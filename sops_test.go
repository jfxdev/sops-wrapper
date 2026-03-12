package sops

import (
	"context"
	"testing"
)

func TestNewCipher(t *testing.T) {
	cipher := NewCipher()
	if cipher == nil {
		t.Fatal("expected non-nil Cipher")
	}
}

func TestDecryptUnsupportedFormat(t *testing.T) {
	cipher := NewCipher()
	ctx := context.Background()

	_, err := cipher.Decrypt(ctx, []byte("some content"), "unsupported-format")
	if err == nil {
		t.Fatal("expected error for unsupported format, got nil")
	}
	expectedErr := "unsupported format: unsupported-format"
	if err.Error() != expectedErr {
		t.Fatalf("expected error '%s', got '%s'", expectedErr, err.Error())
	}
}

func TestEncryptUnsupportedFormat(t *testing.T) {
	cipher := NewCipher()
	ctx := context.Background()

	config := EncryptionConfig{
		Format: "unsupported-format",
	}

	_, err := cipher.Encrypt(ctx, []byte("some content"), config)
	if err == nil {
		t.Fatal("expected error for unsupported format, got nil")
	}
	expectedErr := "unsupported format: unsupported-format"
	if err.Error() != expectedErr {
		t.Fatalf("expected error '%s', got '%s'", expectedErr, err.Error())
	}
}
