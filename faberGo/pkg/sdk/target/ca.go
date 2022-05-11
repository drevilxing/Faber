package target

type TlsCACerts struct {
	Path string `json:"path" yaml:"path"`
}

type TlsCertPair struct {
	Path string `json:"path" yaml:"path"`
	Pem  string `json:"pem" yaml:"pem"`
}

type CAClientConfig struct {
	Key  *TlsCertPair `json:"key" yaml:"key"`
	Cert *TlsCertPair `json:"cert" yaml:"cert"`
}

type CATlsCACerts struct {
	Pem    string          `json:"pem" yaml:"pem"`
	Path   string          `json:"path" yaml:"path"`
	Client *CAClientConfig `json:"client" yaml:"client"`
}

type CARegistrar struct {
	EnrollId     string `json:"enrollId" yaml:"enrollId"`
	EnrollSecret string `json:"enrollSecret" yaml:"enrollSecret"`
}

type CAConfig struct {
	Key        string        `json:"key" yaml:"key"`
	Url        string        `json:"url" yaml:"url"`
	TlsCaCerts *CATlsCACerts `json:"tlsCaCerts" yaml:"tlsCaCerts"`
	Registrar  *CARegistrar  `json:"registrar" yaml:"registrar"`
	CaName     string        `json:"caName" yaml:"caName"`
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
