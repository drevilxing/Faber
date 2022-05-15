package client

type CryptoConfig struct {
	Path string `json:"path" yaml:"path"`
}

func GenerateDefaultCryptoConfig() *CryptoConfig {
	return &CryptoConfig{Path: "/root/opt/crypto-config"} // 路径与生成组织证书过程保持一致。
}
