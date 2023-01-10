# sops-wrapper

Integrated wrapper library for Mozilla SOPS

## Usage

```go
import "github.com/jfxdev/sops-wrapper"
```

```go

cypher := sops.NewCypher()

keys := []entities.EncryptionKey{
    {
        ID:       "arn:aws:kms:us-east-2:XXXXXXXXXXXX:key/YYYYYYYY-YYYY-YYYYY-YYYY-YYYYYYYYYYY",
        Platform: "aws/kms",
        Role:     "arn:aws:iam::XXXXXXXXXXXX:role/your-aws-role",
        Context:  map[string]string{"context": "sops"},
    }
x}

result, err := cypher.Encrypt(body, sops.EncryptionConfig{
		Keys:              keys,
		UnencryptedSuffix: "",
		EncryptedSuffix:   "",
		UnencryptedRegex:  "",
		EncryptedRegex:    "^(data)$",
		ShamirThreshold:   3,
		Format:            "yaml",
})

```