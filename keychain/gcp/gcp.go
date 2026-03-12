package gcp

import (
	"context"

	"github.com/jfxdev/sops-wrapper/keychain/entities"

	"go.mozilla.org/sops/keys"
	"go.mozilla.org/sops/v3/gcpkms"
)

const Alias = "gcp/kms"

func NewKeyGroup(ctx context.Context, key entities.EncryptionKey) (result keys.MasterKey) {
	result = gcpkms.NewMasterKeyFromResourceID(key.ID)
	return
}
