package client

type CryptoConfig struct {
	Path string `json:"path"`
}

func GenerateDefaultCryptoConfig() *CryptoConfig {
	// TODO 需要考虑这一项配置的具体情况，"/root/go"是否为GoPath，并检查是否已经克隆SDK仓库。
	return &CryptoConfig{Path: "/root/go/src/fabric-go-sdk/fixtures/crypto-config"}
}
