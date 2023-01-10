package entities

type EncryptionKey struct {
	ID       string
	Platform string
	Role     string
	Context  map[string]string
}
