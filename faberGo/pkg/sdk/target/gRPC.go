package target

type GrpcOptions struct {
	SslTargetNameOverride string `json:"ssl-target-name-override" yaml:"ssl-target-name-override"`
	KeepAliveTime         string `json:"keep-alive-time" yaml:"keep-alive-time"`
	KeepAliveTimeout      string `json:"keep-alive-timeout" yaml:"keep-alive-timeout"`
	KeepAlivePermit       bool   `json:"keep-alive-permit" yaml:"keep-alive-permit"`
	FailFast              bool   `json:"fail-fast" yaml:"fail-fast"`
	AllowInsecure         bool   `json:"allow-insecure" yaml:"allow-insecure"`
}

func GenerateDefaultGrpcOptions(name string) *GrpcOptions {
	return &GrpcOptions{
		SslTargetNameOverride: name,
		KeepAliveTime:         "5s",
		KeepAliveTimeout:      "6s",
		KeepAlivePermit:       false,
		FailFast:              true,
		AllowInsecure:         false,
	}
}
