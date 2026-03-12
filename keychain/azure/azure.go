package azure

import (
	"context"

	"github.com/jfxdev/sops-wrapper/keychain/entities"
	"go.mozilla.org/sops/keys"
	"go.mozilla.org/sops/v3/azkv"
)

const Alias = "azure/kv"

func NewKeyGroup(ctx context.Context, key entities.EncryptionKey) (result keys.MasterKey) {
	// azkv.NewMasterKeyFromURL returns (MasterKey, error)
	result, _ = azkv.NewMasterKeyFromURL(key.ID)
	return
}
