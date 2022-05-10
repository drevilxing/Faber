package main

import (
	"fmt"
	"net"

	gossh "golang.org/x/crypto/ssh"
)

// 连接信息
type Cli struct {
	user       string
	pwd        string
	addr       string
	client     *gossh.Client
	session    *gossh.Session
	LastResult string
}

// 连接对象
func (c *Cli) Connect() (*Cli, error) {
	config := &gossh.ClientConfig{}
	config.SetDefaults()
	config.User = c.user
	config.Auth = []gossh.AuthMethod{gossh.Password(c.pwd)}
	config.HostKeyCallback = func(hostname string, remote net.Addr, key gossh.PublicKey) error { return nil }
	client, err := gossh.Dial("tcp", c.addr, config)
	if nil != err {
		return c, err
	}
	c.client = client
	return c, nil
}

// 执行shell
func (c Cli) Run(shell string) (string, error) {
	if c.client == nil {
		if _, err := c.Connect(); err != nil {
			return "", err
		}
	}
	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}
	// 关闭会话
	defer session.Close()
	buf, err := session.CombinedOutput(shell)

	c.LastResult = string(buf)
	return c.LastResult, err
}

func BuildEnvironment(s string) {
	cli := Cli{
		addr: "127.0.0.1:22",
		user: "root",
		pwd:  "123456",
	}
	// 建立连接对象
	c, _ := cli.Connect()
	// 退出时关闭连接
	defer c.client.Close()
	c.Run("sudo apt update -y")
	// 安装git docker
	c.Run("sudo apt install git docker.io docker-compose -y")
	c.Run("echo \"{\n\"registry-mirrors\": [\"https://y0qd3iq.mirror.aliyuncs.com\"]\n\"} >> /etc/docker/daemon.json") //这句可能有问题
	c.Run("systemctl restart docker.service")
	// 拉取所需容器
	c.Run("docker pull hyperledger/fabric-tools:2.2.0")
	c.Run("docker pull hyperledger/fabric-ccenv:2.2.0")
	c.Run("docker pull hyperledger/fabric-baseos:2.2.0")
	// 安装go
	c.Run("mkdir develop && cd develop; wget https://studygolang.com/dl/golang/go1.15.6.linux-amd64.tar.gz;")
	c.Run("echo\"\nexport PATH=$PATH:/root/develop/go/bin\nexport GOROOT=/root/develop/go\nexport GOROOT=/root/develop/go\n\" >> ~/.profile;")
	c.Run("source ~/.profile;cd")
	// 拉取节点容器
	c.Run("docker pull hyperledger/fabric-ca:1.4.7")
	c.Run("docker pull hyperledger/fabric-peer:2.2.0")
	c.Run("docker pull hyperledger/fabric-orderer:2.2.0")
	//
	c.Run("mkdir -p go/src/github.com/hyperledger && cd go/src/github.com/hyperledger;git clone https://gitee.com/planewalker/fabric-ca.git")
	if s == "peer" {
		ledger, _ := c.Run("cd fabric-ca;make fabric-ca-client;cp bin/fabric-ca-client /usr/local/bin;chmod 775 /usr/local/bin/fabric-ca-client")
		fmt.Println(ledger)
	} else if s == "orderer" {
		mkdir_cd, _ := c.Run("git clone https://gitee.com/planewalker/fabric.git;cd fabric-ca;cp bin/fabric-ca-client /usr/local/bin;chmod 775 /usr/local/bin/fabric-ca-client;cd;cd fabric;make release;cp release/linux-amd64/bin/configtxgen /usr/local/bin;chmod 775 /usr/local/bin/configtxgen;cd;")
		fmt.Println(mkdir_cd)
	}

}
func main() {
	BuildEnvironment("peer")
}
