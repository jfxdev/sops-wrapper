package vault

import (
	"github.com/jfxdev/sops-wrapper/keychain/entities"

	"go.mozilla.org/sops/keys"
	"go.mozilla.org/sops/v3/hcvault"
)

const Alias = "vault/kms"

func NewKeyGroup(key entities.EncryptionKey) (result keys.MasterKey) {
	result, _ = hcvault.NewMasterKeyFromURI(
		key.ID,
	)
	return
}
