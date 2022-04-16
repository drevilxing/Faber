package target

type TlsCACerts struct {
	Path string `json:"path"`
}

type TlsCertPair struct {
	Path string `json:"path"`
	Pem  string `json:"pem"`
}

type CAClientConfig struct {
	Key  *TlsCertPair `json:"key"`
	Cert *TlsCertPair `json:"cert"`
}

type CATlsCACerts struct {
	Pem    string          `json:"pem"`
	Path   string          `json:"path"`
	Client *CAClientConfig `json:"client"`
}

type CARegistrar struct {
	EnrollId     string `json:"enrollId"`
	EnrollSecret string `json:"enrollSecret"`
}

type CAConfig struct {
	Key        string        `json:"key"`
	Url        string        `json:"url"`
	TlsCaCerts *CATlsCACerts `json:"tlsCaCerts"`
	Registrar  *CARegistrar  `json:"registrar"`
	CaName     string        `json:"caName"`
}

func GenerateDefaultCAConfig(name string, url string) *CAConfig {
	return &CAConfig{
		Key: name,
		Url: "http://localhost:7054",
		TlsCaCerts: &CATlsCACerts{
			Pem:  "ca-cert.pem",
			Path: "/etc/hyperledger/fabric-ca-server/ca-cert.pem",
			// TODO Client相关配置文件还存有疑问，应该是在启动阶段生成有一个对应的用户。
			Client: &CAClientConfig{
				Key: &TlsCertPair{
					Path: "",
					Pem:  "",
				},
				Cert: &TlsCertPair{
					Path: "",
					Pem:  "",
				},
			},
		},
		Registrar: &CARegistrar{
			EnrollId:     "admin",
			EnrollSecret: "adminpw",
		},
		CaName: name,
	}
}
