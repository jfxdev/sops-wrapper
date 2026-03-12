package azure

import (
	"context"
	"testing"

	"github.com/jfxdev/sops-wrapper/keychain/entities"
)

func TestNewKeyGroup(t *testing.T) {
	ctx := context.Background()
	key := entities.EncryptionKey{
		ID: "https://myvault.vault.azure.net/keys/my-key/1a2b3c",
	}

	result := NewKeyGroup(ctx, key)
	if result == nil {
		t.Fatal("expected non-nil MasterKey")
	}
}
