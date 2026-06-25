package utils

import (
	"errors"
	"testing"
)

func TestGenAESKey(t *testing.T) {
	got := GenerateAESKey()

	if len(got.Key) != 16 {
		t.Errorf("Expected length of 16; got %d", len(got.Key))
	}
	if len(got.Iv) != 16 {
		t.Errorf("Expected length of 16; got %d", len(got.Iv))
	}

	if string(got.Iv) == string(got.Key) {
		t.Errorf("Expected number to be random")
	}
}

func TestAES256EncryptDecryptRoundTrip(t *testing.T) {
	key := AESKey{
		Key: []byte("1234567890123456"),
		Iv:  []byte("6543210987654321"),
	}
	plaintext := `{"operation":"read","params":{"device_mac":"default"}}`

	encrypted, err := AES256Encrypt(plaintext, key)
	if err != nil {
		t.Fatalf("AES256Encrypt returned error: %v", err)
	}
	if encrypted == "" || encrypted == plaintext {
		t.Fatalf("expected encrypted value to differ from plaintext; got %q", encrypted)
	}

	decrypted, err := AES256Decrypt(encrypted, key)
	if err != nil {
		t.Fatalf("AES256Decrypt returned error: %v", err)
	}
	if decrypted != plaintext {
		t.Fatalf("expected decrypted plaintext %q; got %q", plaintext, decrypted)
	}
}

func TestAES256EncryptRejectsInvalidKeyLength(t *testing.T) {
	_, err := AES256Encrypt("hello", AESKey{
		Key: []byte("short"),
		Iv:  []byte("6543210987654321"),
	})
	if err == nil {
		t.Fatal("expected invalid key length error")
	}
}

func TestPKCS7PaddingAndUnpadding(t *testing.T) {
	padded, err := pkcs7Padding([]byte("hello"), 8)
	if err != nil {
		t.Fatalf("pkcs7Padding returned error: %v", err)
	}
	if len(padded)%8 != 0 {
		t.Fatalf("expected padded data to align to block size; got %d bytes", len(padded))
	}

	unpadded, err := pkcs7Unpadding(padded, 8)
	if err != nil {
		t.Fatalf("pkcs7Unpadding returned error: %v", err)
	}
	if string(unpadded) != "hello" {
		t.Fatalf("expected original data; got %q", unpadded)
	}
}

func TestPKCS7PaddingRejectsInvalidInput(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		blockSize int
		wantErr   error
	}{
		{
			name:      "invalid block size",
			input:     []byte("hello"),
			blockSize: 0,
			wantErr:   ErrInvalidBlockSize,
		},
		{
			name:      "empty input",
			input:     nil,
			blockSize: 8,
			wantErr:   ErrInvalidPKCS7Data,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := pkcs7Padding(tt.input, tt.blockSize)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected %v; got %v", tt.wantErr, err)
			}
		})
	}
}

func TestPKCS7UnpaddingRejectsInvalidPadding(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		blockSize int
		wantErr   error
	}{
		{
			name:      "invalid block size",
			input:     []byte("12345678"),
			blockSize: 0,
			wantErr:   ErrInvalidBlockSize,
		},
		{
			name:      "empty input",
			input:     nil,
			blockSize: 8,
			wantErr:   ErrInvalidPKCS7Data,
		},
		{
			name:      "not full block",
			input:     []byte("hello"),
			blockSize: 8,
			wantErr:   ErrInvalidPKCS7Padding,
		},
		{
			name:      "wrong padding bytes",
			input:     []byte("abc\x02\x03\x03\x03\x04"),
			blockSize: 8,
			wantErr:   ErrInvalidPKCS7Padding,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := pkcs7Unpadding(tt.input, tt.blockSize)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected %v; got %v", tt.wantErr, err)
			}
		})
	}
}
