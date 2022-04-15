package client

type CryptoConfig struct {
	Path string `json:"path"`
}

func GenerateDefaultCryptoConfig() *CryptoConfig {
	// TODO 需要考虑这一项配置的具体情况
	return &CryptoConfig{Path: "/etc/hyperledger/crypto-config"}
}
