# sops-wrapper

Integrated library wrapper for [**Mozilla SOPS**](https://github.com/mozilla/sops), ideal for run encryption/decryption on apps and web services, without need to install the official binary.

## Usage

### Simple Encryption Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/jfxdev/sops-wrapper"
	"github.com/jfxdev/sops-wrapper/keychain/entities"
)

func main() {
    ctx := context.Background()
	cipher := sops.NewCipher()

	config := sops.EncryptionConfig{
		Format: sops.FormatYAML,
		Keys: []entities.EncryptionKey{
			{
				Platform: "aws/kms",
				ID:       "arn:aws:kms:us-east-1:1234567890:key/abc-123",
			},
		},
	}

	encrypted, err := cipher.Encrypt(ctx, []byte(`{"secret": "value"}`), config)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(encrypted))
}
```

### Supported Keychains

Here are examples of how to populate `EncryptionConfig.Keys` for each supported cloud provider:

**1. AWS KMS**
```go
entities.EncryptionKey{
    Platform: "aws/kms",
    ID:       "arn:aws:kms:us-east-1:1234567890:key/abc-123",
    Role:     "",                  // Optional generic IAM Role ARN
    Context:  "user:api,env:prod", // Optional Encryption Context limits
}
```

**2. Google Cloud KMS**
```go
entities.EncryptionKey{
    Platform: "gcp/kms",
    ID:       "projects/my-project/locations/global/keyRings/my-ring/cryptoKeys/my-key",
}
```

**3. Azure Key Vault (AKV)**
```go
entities.EncryptionKey{
    Platform: "azure/kv",
    ID:       "https://myvault.vault.azure.net/keys/my-key/1a2b3c",
}
```

**4. HashiCorp Vault**
```go
entities.EncryptionKey{
    Platform: "vault/kms",
    Parameters: map[string]string{
        "url":         "https://vault.corp.local:8200",
        "engine_path": "sops",
        "key_path":    "my-encryption-key",
    },
}
```

## Key Rotation

Rotating keys involves decrypting existing documents and re-encrypting them with a new set of keys (generating a fresh DEK). To do this seamlessly:

```go
newConfig := sops.EncryptionConfig{
    Format: sops.FormatYAML,
    Keys: []entities.EncryptionKey{ /* New Keys here */ },
}

rotatedContent, err := cipher.Rotate(ctx, encryptedPayloadBytes, newConfig)
```