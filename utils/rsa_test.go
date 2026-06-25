package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"math/big"
	"testing"
)

func TestGenerateRsaKey(t *testing.T) {
	got, err := GenerateRsaKey([]string{"c34f", "10001"})
	if err != nil {
		t.Fatalf("GenerateRsaKey returned error: %v", err)
	}

	if got.N.Cmp(big.NewInt(0xc34f)) != 0 {
		t.Fatalf("unexpected modulus: %s", got.N.Text(16))
	}
	if got.E != 65537 {
		t.Fatalf("unexpected exponent: %d", got.E)
	}
}

func TestGenerateRsaKeyRejectsInvalidExponent(t *testing.T) {
	_, err := GenerateRsaKey([]string{"c34f", "not-hex"})
	if err == nil {
		t.Fatal("expected error for invalid exponent")
	}
}

func TestEncryptRsaCanBeDecrypted(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatalf("failed to generate RSA key: %v", err)
	}

	encrypted, err := EncryptRsa("hello", &privateKey.PublicKey)
	if err != nil {
		t.Fatalf("EncryptRsa returned error: %v", err)
	}

	ciphertext, err := hex.DecodeString(encrypted)
	if err != nil {
		t.Fatalf("encrypted value is not valid hex: %v", err)
	}
	if len(ciphertext) != privateKey.Size() {
		t.Fatalf("expected ciphertext size %d; got %d", privateKey.Size(), len(ciphertext))
	}

	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		t.Fatalf("failed to decrypt ciphertext: %v", err)
	}
	if string(plaintext) != "hello" {
		t.Fatalf("expected plaintext hello; got %q", plaintext)
	}
}

func TestEncryptRsaRejectsInvalidKey(t *testing.T) {
	_, err := EncryptRsa("hello", nil)
	if err == nil {
		t.Fatal("expected error for nil key")
	}
}
