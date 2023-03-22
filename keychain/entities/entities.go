package entities

type EncryptionKey struct {
	ID         string
	Platform   string
	Role       string
	Parameters map[string]string
	Context    map[string]string
}
