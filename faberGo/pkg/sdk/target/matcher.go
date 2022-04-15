package target

type Matcher struct {
	Pattern                             string `json:"pattern"`
	UrlSubstitutionExp                  string `json:"urlSubstitutionExp"`
	SslTargetOverrideUrlSubstitutionExp string `json:"sslTargetOverrideUrlSubstitutionExp"`
	MappedHost                          string `json:"mappedHost"`
	IgnoreEndpoint                      bool   `json:"ignoreEndpoint"`
}

type EntryMatcher struct {
	Peers                *[]*Matcher `json:"peers"`
	Orderer              *[]*Matcher `json:"orderer"`
	CertificateAuthority *[]*Matcher `json:"certificateAuthority"`
	Channel              *[]*Matcher `json:"channel"`
}

func GenerateDefaultEntryMatcher() *EntryMatcher {
	return &EntryMatcher{
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

func (that *EntryMatcher) AddPeer(key string) {
	for _, element := range *that.Peers {
		if key == element.MappedHost {
			return
		}
	}
	*that.Peers = append(*that.Peers, GenerateMatcherCommon("(\\w*)"+key+"(\\w*)", "localhost:7051", "localhost:7053", key))
}

func (that *EntryMatcher) AddCA(key string) {
	*that.CertificateAuthority = append(*that.CertificateAuthority, GenerateMatcherCommon("(\\w*)"+key+"(\\w*)", "http://localhost:7054", "", key))
}

func (that *EntryMatcher) AddOrderer(key string) {
	*that.Orderer = append(*that.Orderer, GenerateMatcherCommon("(\\w*)"+key+"(\\w*)", "localhost:7050", key, key))

}
