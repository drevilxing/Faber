package client

type SecurityDefault struct {
	Provider      string `json:"provider"`
	HashAlgorithm string `json:"hashAlgorithm"`
	SoftVerify    bool   `json:"softVerify"`
	Level         int64  `json:"level"`
	Pin           string `json:"pin"`
	Label         string `json:"label"`
	Library       string `json:"library"`
}

type Security struct {
	Enable  bool             `json:"enable"`
	Default *SecurityDefault `json:"default"`
}

type BCCSP struct {
	Security *Security `json:"security"`
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
