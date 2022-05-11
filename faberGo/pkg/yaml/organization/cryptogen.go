package organization

import (
	"gopkg.in/yaml.v2"
	"os"
)

// Orderer Organizations.

type Specs struct {
	Hostname string    `yaml:"Hostname"`
	SANS     *[]string `yaml:"SANS"`
}

type OrdererOrg struct {
	Name          string    `yaml:"Name"`
	Domain        string    `yaml:"Domain"`
	EnableNodeOUs bool      `yaml:"EnableNodeOUs"`
	Specs         *[]*Specs `yaml:"Specs"`
}

// Peer Organizations.

type Users struct {
	Count int64 `yaml:"Count"`
}

type Template struct {
	Count int64     `yaml:"Count"`
	SANS  *[]string `yaml:"SANS"`
}

type PeerOrg struct {
	Name          string    `yaml:"Name"`
	Domain        string    `yaml:"Domain"`
	EnableNodeOUs bool      `yaml:"EnableNodeOUs"`
	Template      *Template `yaml:"Template"`
	Users         *Users    `yaml:"Users"`
}

// Cryptogen Config.

type CryptogenConfig struct {
	OrdererOrgs *[]*OrdererOrg `yaml:"OrdererOrgs"`
	PeerOrgs    *[]*PeerOrg    `yaml:"PeerOrgs"`
}

func GenerateEmptyCryptogenConfig() *CryptogenConfig {
	return &CryptogenConfig{
		OrdererOrgs: &[]*OrdererOrg{},
		PeerOrgs:    &[]*PeerOrg{},
	}
}

func (that *CryptogenConfig) AddOrdererOrg(name, domain, hostname string, sans []string) {
	*that.OrdererOrgs = append(*that.OrdererOrgs, &OrdererOrg{
		Name:          name,
		Domain:        domain,
		EnableNodeOUs: true,
		Specs: &[]*Specs{
			&Specs{
				Hostname: hostname,
				SANS:     &sans,
			},
		},
	})
}

func (that *CryptogenConfig) AddPeerOrg(name, domain string, sans []string, users int64) {
	*that.PeerOrgs = append(*that.PeerOrgs, &PeerOrg{
		Name:          name,
		Domain:        domain,
		EnableNodeOUs: true,
		Template: &Template{
			Count: int64(len(sans)),
			SANS:  &sans,
		},
		Users: &Users{
			Count: users,
		},
	})
}

func (that *CryptogenConfig) GenerateYaml(path string) error {
	data, err := yaml.Marshal(*that)
	if nil != err {
		return err
	}
	file, err := os.OpenFile(path+"/crypto-config.yaml", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	if nil != err {
		return err
	}
	_, err = file.Write(data)
	if nil != err {
		return err
	}
	return file.Sync()
}
