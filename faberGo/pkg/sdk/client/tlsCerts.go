package client

type TlsKey struct {
	Path string `json:"path" yaml:"path"`
}

type TlsCert struct {
	Path string `json:"path" yaml:"path"`
}

type TlsCerts struct {
	Key            *TlsKey  `json:"key" yaml:"key"`
	Cert           *TlsCert `json:"cert" yaml:"cert"`
	SystemCertPool bool     `json:"systemCertPool" yaml:"systemCertPool"`
}

func GenerateDefaultTlsCerts() *TlsCerts {
	return &TlsCerts{
		Key:            &TlsKey{Path: ""}, // 默认使用Orderer组织的Admin用户
		Cert:           &TlsCert{Path: ""},
		SystemCertPool: true,
	}
}
