package target

type Matcher struct {
	Pattern                             string `json:"pattern" yaml:"pattern"`
	UrlSubstitutionExp                  string `json:"urlSubstitutionExp" yaml:"urlSubstitutionExp"`
	SslTargetOverrideUrlSubstitutionExp string `json:"sslTargetOverrideUrlSubstitutionExp" yaml:"sslTargetOverrideUrlSubstitutionExp"`
	MappedHost                          string `json:"mappedHost" yaml:"mappedHost"`
	IgnoreEndpoint                      bool   `json:"ignoreEndpoint" yaml:"ignoreEndpoint"`
}

type EntityMatcher struct {
	Peers                *[]*Matcher `json:"peers" yaml:"peers"`
	Orderer              *[]*Matcher `json:"orderer" yaml:"orderer"`
	CertificateAuthority *[]*Matcher `json:"certificateAuthority" yaml:"certificateAuthority"`
	Channel              *[]*Matcher `json:"channel" yaml:"channel"`
}

func GenerateDefaultEntityMatcher() *EntityMatcher {
	return &EntityMatcher{
		Peers:                &[]*Matcher{},
		Orderer:              &[]*Matcher{},
		CertificateAuthority: &[]*Matcher{},
		Channel:              &[]*Matcher{},
	}
}

func GenerateMatcherCommon(pattern string, urlSubstitutionExp string, sslTargetOverrideUrlSubstitutionExp string, mappedHost string) *Matcher {
	return &Matcher{
		Pattern:                             pattern,
		UrlSubstitutionExp:                  urlSubstitutionExp,
		SslTargetOverrideUrlSubstitutionExp: sslTargetOverrideUrlSubstitutionExp,
		MappedHost:                          mappedHost,
	}
}

func GenerateMatcherIgnoreEndpoint(pattern string) *Matcher {
	return &Matcher{
		Pattern:        pattern,
		IgnoreEndpoint: true,
	}
}

func (that *EntityMatcher) AddPeer(key string) {
	for _, element := range *that.Peers {
		if key == element.MappedHost {
			return
		}
	}
	*that.Peers = append(*that.Peers, GenerateMatcherCommon("(\\w*)"+key+"(\\w*)", "localhost:7051", "localhost:7053", key))
}

func (that *EntityMatcher) AddCA(key string) {
	*that.CertificateAuthority = append(*that.CertificateAuthority, GenerateMatcherCommon("(\\w*)"+key+"(\\w*)", "http://localhost:7054", "", key))
}

func (that *EntityMatcher) AddOrderer(key string) {
	*that.Orderer = append(*that.Orderer, GenerateMatcherCommon("(\\w*)"+key+"(\\w*)", "localhost:7050", key, key))

}
