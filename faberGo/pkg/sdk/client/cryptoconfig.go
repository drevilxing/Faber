package client

type CryptoConfig struct {
	Path string `json:"path"`
}

func GenerateDefaultCryptoConfig() *CryptoConfig {
	return &CryptoConfig{Path: ""}
}
