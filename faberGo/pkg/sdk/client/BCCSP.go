package client

type SecurityDefault struct {
	Provider      string `json:"provider" yaml:"provider"`
	HashAlgorithm string `json:"hashAlgorithm" yaml:"hashAlgorithm"`
	SoftVerify    bool   `json:"softVerify" yaml:"softVerify"`
	Level         int64  `json:"level" yaml:"level"`
	Pin           string `json:"pin" yaml:"pin"`
	Label         string `json:"label" yaml:"label"`
	Library       string `json:"library" yaml:"library"`
}

type Security struct {
	Enable  bool             `json:"enable" yaml:"enable"`
	Default *SecurityDefault `json:"default" yaml:"default"`
}

type BCCSP struct {
	Security *Security `json:"security" yaml:"security"`
}

func GenerateDefaultBCCSP() *BCCSP {
	return &BCCSP{Security: &Security{
		Enable: true,
		Default: &SecurityDefault{
			Provider:      "SW",
			HashAlgorithm: "SHA2",
			SoftVerify:    true,
			Level:         256,
			Pin:           "pin",
			Label:         "label",
			Library:       "",
		},
	}}
}
