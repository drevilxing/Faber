package client

type CryptoStore struct {
	Path string `json:"path"`
}

type CredentialStore struct {
	Path        string       `json:"path" yaml:"path"`
	CryptoStore *CryptoStore `json:"cryptoStore" yaml:"cryptoStore"`
}

func GenerateDefaultCredentialStore() *CredentialStore {
	// TODO 这一项配置文件需要看情况，/etc/hyperledger/sdk/$path
	return &CredentialStore{
		Path:        "/tmp/state-store",
		CryptoStore: &CryptoStore{Path: "/tmp/msp"},
	}
}
