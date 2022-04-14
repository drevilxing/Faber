package client

type CryptoStore struct {
	Path string `json:"path"`
}

type CredentialStore struct {
	Path        string       `json:"path"`
	CryptoStore *CryptoStore `json:"cryptoStore"`
}

func GenerateDefaultCredentialStore() *CredentialStore {
	return &CredentialStore{
		Path:        "",
		CryptoStore: &CryptoStore{Path: ""},
	}
}
