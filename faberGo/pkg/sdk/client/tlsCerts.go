package client

type TlsKey struct {
	Path string `json:"path"`
}

type TlsCert struct {
	Path string `json:"path"`
}

type TlsCerts struct {
	Key  *TlsKey  `json:"key"`
	Cert *TlsCert `json:"cert"`
}

func GenerateDefaultTlsCerts() *TlsCerts {
	return &TlsCerts{
		Key:  &TlsKey{Path: ""},
		Cert: &TlsCert{Path: ""},
	}
}
