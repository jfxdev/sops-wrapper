package vault

import (
	"context"
	"testing"

	"github.com/jfxdev/sops-wrapper/keychain/entities"
)

func TestNewKeyGroup(t *testing.T) {
	ctx := context.Background()
	key := entities.EncryptionKey{
		Parameters: map[string]string{
			"url":         "https://vault.corp.local:8200",
			"engine_path": "sops",
			"key_path":    "my-key",
		},
	}

	result := NewKeyGroup(ctx, key)
	if result == nil {
		t.Fatal("expected non-nil MasterKey")
	}
}
