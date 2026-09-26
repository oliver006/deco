package deco

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/oliver006/deco/utils"
)

func TestAuthenticationErrors(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatal(err)
	}
	aes := &utils.AESKey{Key: []byte("1234567890123456"), Iv: []byte("6543210987654321")}
	badJSON, err := utils.AES256Encrypt(`{"result":`, *aes)
	if err != nil {
		t.Fatal(err)
	}
	badJSONResponse, err := json.Marshal(map[string]string{"data": badJSON})
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		name   string
		status int
		body   string
		err    error
		auth   bool
	}{
		{name: "expired session", status: 200, body: `{"data":""}`, auth: true},
		{name: "unauthorized", status: 401, auth: true},
		{name: "forbidden", status: 403, auth: true},
		{name: "timeout", err: context.DeadlineExceeded},
		{name: "server error", status: 500},
		{name: "missing data", status: 200, body: `{}`},
		{name: "null data", status: 200, body: `{"data":null}`},
		{name: "empty HTTP body", status: 200},
		{name: "malformed envelope", status: 200, body: `{"data":`},
		{name: "malformed decrypted JSON", status: 200, body: string(badJSONResponse)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			client := &Client{
				aes: aes, rsa: &key.PublicKey, hash: "test", stok: "test-token",
				c: &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
					if tc.err != nil {
						return nil, tc.err
					}
					return &http.Response{StatusCode: tc.status, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(tc.body))}, nil
				})},
			}
			_, err := client.Performance()
			if err == nil {
				t.Fatal("expected request failure")
			}
			if got := errors.Is(err, ErrAuthentication); got != tc.auth {
				t.Fatalf("authentication error = %t, want %t: %v", got, tc.auth, err)
			}
			if tc.err != nil && !errors.Is(err, tc.err) {
				t.Fatalf("request error lost: %v", err)
			}
		})
	}
}
