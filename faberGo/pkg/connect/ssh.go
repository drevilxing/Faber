package connect

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
)

type Commands struct {
	Commands *[]string
	Device   *Device
}

func (that *Commands) Append(command string) {
	*that.Commands = append(*that.Commands, command)
}

func (that *Commands) Execute() (*[]string, *[]error, error) {
	results := &[]string{}
	errs := &[]error{}
	err := that.Device.Connect()
	if nil != err {
		return nil, nil, err
	}
	for i := range *that.Commands {
		resIn, errIn := that.Device.Execute((*that.Commands)[i])
		*results = append(*results, resIn)
		*errs = append(*errs, errIn)
	}
	return results, errs, nil
}

type Device struct {
	IP       string
	Port     string
	User     string
	Password string
	Key      string
	UsePwd   bool
	client   *ssh.Client
}

func (that *Device) Connect() (err error) {
	config := &ssh.ClientConfig{}
	config.SetDefaults()
	config.User = that.User
	if that.UsePwd {
		config.Auth = []ssh.AuthMethod{
			ssh.Password(that.Password),
		}
	} else {
		config.Auth = []ssh.AuthMethod{}
	}
	config.HostKeyCallback = func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	that.client, err = ssh.Dial("tcp", that.IP+":"+that.Port, config)
	return err
}

func (that *Device) Close() error {
	return that.client.Close()
}

func (that *Device) Execute(command string) (result string, err error) {
	if nil == that.client {
		err = that.Connect()
		if nil != err {
			return "", err
		}
	}
	session, err := that.client.NewSession()
	if nil != err {
		return "", err
	}
	defer func(session *ssh.Session) {
		errIn := session.Close()
		if errIn != nil {
			fmt.Println(errIn.Error())
			return
		}
	}(session)
	buffer, err := session.CombinedOutput(command)
	if nil != err {
		return "", err
	}
	result = string(buffer)
	return result, nil
}
