package deco

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/oliver006/deco/utils"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestEndpointArgsQueryParams(t *testing.T) {
	args := EndpointArgs{form: "performance"}
	got := args.queryParams()

	if got.Get("form") != "performance" {
		t.Fatalf("expected form=performance; got %q", got.Get("form"))
	}
}

func TestDoPostSendsRequestAndDecodesResponse(t *testing.T) {
	oldBaseURL := baseURL
	baseURL = url.URL{
		Scheme: "http",
		Host:   "deco.local",
		Path:   "/cgi-bin/luci/",
	}
	t.Cleanup(func() {
		baseURL = oldBaseURL
	})

	client := &Client{
		c: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				if req.Method != http.MethodPost {
					t.Errorf("expected POST request; got %s", req.Method)
				}
				if req.URL.Host != "deco.local" {
					t.Errorf("expected host deco.local; got %s", req.URL.Host)
				}
				if !strings.HasSuffix(req.URL.Path, "/login") {
					t.Errorf("expected login path; got %s", req.URL.Path)
				}
				if req.URL.Query().Get("form") != "keys" {
					t.Errorf("expected form=keys; got %q", req.URL.Query().Get("form"))
				}
				if got := req.Header.Get("Content-Type"); got != "application/json" {
					t.Errorf("expected application/json content type; got %q", got)
				}

				body, err := io.ReadAll(req.Body)
				if err != nil {
					t.Fatalf("failed to read request body: %v", err)
				}
				if string(body) != string(readBody) {
					t.Errorf("unexpected body: %s", body)
				}

				return &http.Response{
					StatusCode: http.StatusOK,
					Header:     make(http.Header),
					Body:       io.NopCloser(strings.NewReader(`{"ok":true}`)),
				}, nil
			}),
		},
	}

	var got struct {
		OK bool `json:"ok"`
	}
	if err := client.doPost(";stok=/login", EndpointArgs{form: "keys"}, readBody, &got); err != nil {
		t.Fatalf("doPost returned error: %v", err)
	}
	if !got.OK {
		t.Fatal("expected response to be decoded")
	}
}

func TestDoPostReturnsErrorForUnexpectedStatus(t *testing.T) {
	client := &Client{
		c: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusTeapot,
					Header:     make(http.Header),
					Body:       io.NopCloser(strings.NewReader(`{}`)),
				}, nil
			}),
		},
	}

	var got map[string]interface{}
	err := client.doPost(";stok=/login", EndpointArgs{form: "keys"}, readBody, &got)
	if err == nil {
		t.Fatal("expected error for non-OK response")
	}
	if !strings.Contains(err.Error(), "unexpected status code: 418") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestClientListForDeviceSendsEncryptedRequestAndDecodesNames(t *testing.T) {
	oldBaseURL := baseURL
	baseURL = url.URL{
		Scheme: "http",
		Host:   "deco.local",
		Path:   "/cgi-bin/luci/",
	}
	t.Cleanup(func() {
		baseURL = oldBaseURL
	})

	aesKey := &utils.AESKey{
		Key: []byte("1234567890123456"),
		Iv:  []byte("6543210987654321"),
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatalf("failed to generate RSA key: %v", err)
	}

	client := &Client{
		aes:      aesKey,
		rsa:      &privateKey.PublicKey,
		hash:     "hash",
		stok:     "token",
		sequence: 7,
		c: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				if !strings.Contains(req.URL.Path, "/admin/client") {
					t.Errorf("expected client endpoint; got %s", req.URL.Path)
				}
				if req.URL.Query().Get("form") != "client_list" {
					t.Errorf("expected form=client_list; got %q", req.URL.Query().Get("form"))
				}

				body, err := io.ReadAll(req.Body)
				if err != nil {
					t.Fatalf("failed to read request body: %v", err)
				}
				form, err := url.ParseQuery(string(body))
				if err != nil {
					t.Fatalf("failed to parse encrypted request body: %v", err)
				}
				if form.Get("sign") == "" {
					t.Error("expected signed request")
				}

				decrypted, err := utils.AES256Decrypt(form.Get("data"), *aesKey)
				if err != nil {
					t.Fatalf("failed to decrypt request data: %v", err)
				}
				var request request
				if err := json.Unmarshal([]byte(decrypted), &request); err != nil {
					t.Fatalf("failed to decode request data: %v", err)
				}
				if request.Operation != "read" {
					t.Errorf("expected read operation; got %q", request.Operation)
				}
				if request.Params["device_mac"] != "default" {
					t.Errorf("expected default device MAC; got %#v", request.Params["device_mac"])
				}

				payload := `{"error_code":0,"result":{"client_list":[{"name":"` + base64.StdEncoding.EncodeToString([]byte("Office Laptop")) + `","online":true}]}}`
				encrypted, err := utils.AES256Encrypt(payload, *aesKey)
				if err != nil {
					t.Fatalf("failed to encrypt response: %v", err)
				}
				responseBody, err := json.Marshal(response{Data: encrypted})
				if err != nil {
					t.Fatalf("failed to marshal response: %v", err)
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Header:     make(http.Header),
					Body:       io.NopCloser(strings.NewReader(string(responseBody))),
				}, nil
			}),
		},
	}

	got, err := client.ClientListForDevice("")
	if err != nil {
		t.Fatalf("ClientListForDevice returned error: %v", err)
	}
	if len(got.Result.ClientList) != 1 {
		t.Fatalf("expected one client; got %d", len(got.Result.ClientList))
	}
	if got.Result.ClientList[0].Name != "Office Laptop" {
		t.Fatalf("expected decoded client name; got %q", got.Result.ClientList[0].Name)
	}
}
