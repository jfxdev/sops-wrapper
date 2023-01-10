package aws

import (
	"fmt"

	"github.com/jfxdev/sops-wrapper/keychain/entities"

	"go.mozilla.org/sops/keys"
	"go.mozilla.org/sops/v3/kms"
)

const Alias = "aws/kms"

func NewKeyGroup(key entities.EncryptionKey) (result keys.MasterKey) {
	result = kms.NewMasterKeyFromArn(
		fmt.Sprintf(
			"%s+%s",
			key.ID,
			key.Role,
		),
		kms.ParseKMSContext(key.Context),
		"",
	)
	return
}
