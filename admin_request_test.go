package deco

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/oliver006/deco/utils"
)

type adminMethodCase struct {
	name      string
	path      string
	form      string
	operation string
	call      func(*Client, map[string]interface{}) (map[string]interface{}, error)
}

func testAdminMethodCases(t *testing.T, cases []adminMethodCase) {
	t.Helper()

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

	requestIndex := 0
	client := &Client{
		aes:      aesKey,
		rsa:      &privateKey.PublicKey,
		hash:     "hash",
		stok:     "token",
		sequence: 7,
		c: &http.Client{
			Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				if requestIndex >= len(cases) {
					t.Fatalf("unexpected extra request to %s", req.URL.String())
				}
				expected := cases[requestIndex]
				requestIndex++

				if req.Method != http.MethodPost {
					t.Errorf("%s: expected POST request; got %s", expected.name, req.Method)
				}
				if !strings.HasSuffix(req.URL.Path, ";stok=token"+expected.path) {
					t.Errorf("%s: expected path suffix %q; got %q", expected.name, ";stok=token"+expected.path, req.URL.Path)
				}

				query := req.URL.Query()
				_, hasForm := query["form"]
				if expected.form == "" && hasForm {
					t.Errorf("%s: expected no form query; got %q", expected.name, query.Encode())
				}
				if expected.form != "" && query.Get("form") != expected.form {
					t.Errorf("%s: expected form=%s; got %q", expected.name, expected.form, query.Get("form"))
				}

				body, err := io.ReadAll(req.Body)
				if err != nil {
					t.Fatalf("%s: failed to read request body: %v", expected.name, err)
				}
				formBody, err := url.ParseQuery(string(body))
				if err != nil {
					t.Fatalf("%s: failed to parse encrypted request body: %v", expected.name, err)
				}
				if formBody.Get("sign") == "" {
					t.Errorf("%s: expected signed request", expected.name)
				}

				decrypted, err := utils.AES256Decrypt(formBody.Get("data"), *aesKey)
				if err != nil {
					t.Fatalf("%s: failed to decrypt request data: %v", expected.name, err)
				}
				var got request
				if err := json.Unmarshal([]byte(decrypted), &got); err != nil {
					t.Fatalf("%s: failed to decode request data: %v", expected.name, err)
				}
				if got.Operation != expected.operation {
					t.Errorf("%s: expected operation %q; got %q", expected.name, expected.operation, got.Operation)
				}
				if got.Params["case"] != expected.name {
					t.Errorf("%s: expected params.case=%q; got %#v", expected.name, expected.name, got.Params["case"])
				}

				payload, err := json.Marshal(map[string]interface{}{
					"error_code": 0,
					"result": map[string]interface{}{
						"method": expected.name,
					},
				})
				if err != nil {
					t.Fatalf("%s: failed to marshal response payload: %v", expected.name, err)
				}
				encrypted, err := utils.AES256Encrypt(string(payload), *aesKey)
				if err != nil {
					t.Fatalf("%s: failed to encrypt response: %v", expected.name, err)
				}
				responseBody, err := json.Marshal(response{Data: &encrypted})
				if err != nil {
					t.Fatalf("%s: failed to marshal response: %v", expected.name, err)
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Header:     make(http.Header),
					Body:       io.NopCloser(strings.NewReader(string(responseBody))),
				}, nil
			}),
		},
	}

	for _, tc := range cases {
		got, err := tc.call(client, map[string]interface{}{"case": tc.name})
		if err != nil {
			t.Fatalf("%s: returned error: %v", tc.name, err)
		}
		if got["error_code"] != float64(0) {
			t.Fatalf("%s: expected error_code 0; got %#v", tc.name, got["error_code"])
		}
	}
	if requestIndex != len(cases) {
		t.Fatalf("expected %d requests; got %d", len(cases), requestIndex)
	}
}
