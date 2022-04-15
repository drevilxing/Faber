package client

type CryptoStore struct {
	Path string `json:"path"`
}

type CredentialStore struct {
	Path        string       `json:"path"`
	CryptoStore *CryptoStore `json:"cryptoStore"`
}

func GenerateDefaultCredentialStore() *CredentialStore {
	// TODO 这一项配置文件需要看情况
	return &CredentialStore{
		Path:        "/etc/hyperledger/sdk/store",
		CryptoStore: &CryptoStore{Path: "/etc/hyperledger/sdk/msp"},
	}
}
