package main

import (
	"bytes"
	"faberGo/pkg/connect"
	"fmt"
	"os/exec"
)

const IP = "192.168.3.20"
const Password = "linux"
const User = "root"

func Example() {
	commands := &connect.Commands{
		Commands: &[]string{},
		Device: &connect.Device{
			IP:       IP,
			Port:     "22",
			User:     User,
			Password: Password,
			Key:      "",
			UsePwd:   true,
		},
	}
	//Environment(commands)
	//Organization(commands)
	Peer(commands)

	results, errs, err := commands.Execute()
	if nil != err {
		fmt.Println(err.Error())
	}
	for i := range *results {
		if nil != (*errs)[i] {
			fmt.Println("err:", err.Error())
		} else {
			fmt.Println((*results)[i])
		}
	}
}

func Environment(commands *connect.Commands) {
	cli := commands.Device

	err := cli.ExecuteAndPrint("apt-get update -y")
	if nil != err {
		fmt.Println("err:", err)
	}

	err = cli.ExecuteAndPrint("apt-get install apt-transport-https ca-certificates -y")
	if nil != err {
		fmt.Println("err:", err)
	}

	err = cli.ExecuteAndPrint("mv /etc/apt/sources.list /etc/apt/sources.list.bak")
	if nil != err {
		fmt.Println("err:", err)
	}

	err = cli.ExecuteAndPrint("echo \"deb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye main contrib non-free\n# deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye main contrib non-free\ndeb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-updates main contrib non-free\n# deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-updates main contrib non-free\n\ndeb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-backports main contrib non-free\n# deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-backports main contrib non-free\n\ndeb https://mirrors.tuna.tsinghua.edu.cn/debian-security bullseye-security main contrib non-free\n# deb-src https://mirrors.tuna.tsinghua.edu.cn/debian-security bullseye-security main contrib non-free\" >> /etc/apt/sources.list\n")
	if nil != err {
		fmt.Println("err:", err)
	}

	err = cli.ExecuteAndPrint("apt-get update -y")
	if nil != err {
		fmt.Println("err:", err)
	}

	err = cli.ExecuteAndPrint("apt-get install vim wget git docker docker.io docker-compose -y")
	if nil != err {
		fmt.Println("err:", err)
	}

	err = cli.ExecuteAndPrint("git clone https://gitee.com/EternallyAscend/Faber")
	if nil != err {
		fmt.Println("err:", err)
	}

	//err = cli.ExecuteAndPrint("wget https://studygolang.com/dl/golang/go1.17.3.linux-amd64.tar.gz")
	//if nil != err {
	//	fmt.Println("err:", err)
	//}

	err = cli.ExecuteAndPrint("tar -C /usr/bin -xzf go1.17.3.linux-amd64.tar.gz")
	if nil != err {
		fmt.Println("err:", err)
	}

	err = cli.ExecuteAndPrint("echo \"export PATH=$PATH:/usr/bin/go/bin\" >> ~/.profile")
	if nil != err {
		fmt.Println("err:", err)
	}

	err = cli.ExecuteAndPrint("source ~/.profile")
	if nil != err {
		fmt.Println("err:", err)
	}

	//err = cli.ExecuteAndPrint("git clone -b binary-1.2.2 https://gitee.com/Ambuland/packages.git")
	//if nil != err {
	//	fmt.Println("err:", err)
	//}

	err = cli.ExecuteAndPrint("cd packages")
	if nil != err {
		fmt.Println("err:", err)
	}

	err = cli.ExecuteAndPrint("tar -xzvf hyperledger-fabric-linux-amd64-2.2.0.tar.gz -C /usr")
	if nil != err {
		fmt.Println("err:", err)
	}

	//err = cli.ExecuteAndPrint("cp ./bin/* /bin && cp ./bin/* /usr/bin")
	//if nil != err {
	//	fmt.Println("err:", err)
	//}

	err = cli.ExecuteAndPrint("tar -xzvf hyperledger-fabric-ca-linux-amd64-1.4.8.tar.gz -C /usr")
	if nil != err {
		fmt.Println("err:", err)
	}

	//err = cli.ExecuteAndPrint("cp ./bin/* /bin && cp ./bin/* /usr/bin")
	//if nil != err {
	//	fmt.Println("err:", err)
	//}
	//
	//err = cli.ExecuteAndPrint("cp ./config/* /bin && cp ./config/* /usr/bin")
	//if nil != err {
	//	fmt.Println("err:", err)
	//}

	err = cli.ExecuteAndPrint("cd /root/Faber && chmod +x ./installGo.sh && ./installGo.sh")
	if nil != err {
		fmt.Println("err:", err)
	}

	//err = cli.ExecuteAndPrint("export PATH=$PATH:/usr/bin/go/bin")
	//if nil != err {
	//	fmt.Println("err:", err)
	//}
	//
	//err = cli.ExecuteAndPrint("source ~/.profile")
	//if nil != err {
	//	fmt.Println("err:", err)
	//}

	//err = cli.ExecuteAndPrint("go version")
	//if nil != err {
	//	fmt.Println("err:", err)
	//}
	//
	//err = cli.ExecuteAndPrint("go env -w GO111MODULE=on")
	//if nil != err {
	//	fmt.Println("err:", err)
	//}
	//
	//err = cli.ExecuteAndPrint("go env -w GOPROXY=https://goproxy.cn,direct")
	//if nil != err {
	//	fmt.Println("err:", err)
	//}

	err = cli.ExecuteAndPrint("cd")
	if nil != err {
		fmt.Println("err:", err)
	}

	err = cli.ExecuteAndPrint("docker pull hyperledger/fabric-peer:2.2.0")
	if nil != err {
		fmt.Println("err:", err)
	}

	err = cli.ExecuteAndPrint("docker pull hyperledger/fabric-orderer:2.2.0")
	if nil != err {
		fmt.Println("err:", err)
	}

	err = cli.ExecuteAndPrint("docker pull hyperledger/fabric-ca:1.4.7")
	if nil != err {
		fmt.Println("err:", err)
	}
}

func Organization(commands *connect.Commands) {
	cli := commands.Device

	err := cli.ExecuteAndPrint("mkdir /root/opt")
	if nil != err {
		fmt.Println("err:", err)
	}

	err = cli.ExecuteAndPrint("cp -r /root/Faber/crypto-config /root/opt")
	if nil != err {
		fmt.Println("err:", err)
	}

	err = cli.ExecuteAndPrint("cryptogen generate --config=/root/opt/crypto-config/crypto-config.yaml --output=\"crypto-config\"")
	if nil != err {
		fmt.Println("err:", err)
	}
}

func Peer(commands *connect.Commands) {
	cli := commands.Device

	err := cli.ExecuteAndPrint("cd /root/Faber/fabric-draw-back")
	if nil != err {
		fmt.Println("err:", err)
	}
	//err = cli.ExecuteAndPrint("python3 node_build.py")
	//if nil != err {
	//	fmt.Println("err:", err)
	//}

	cmd := exec.Command("/bin/bash", "-c", "cd fabric-draw-back && python3 node_generator.py")

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err = cmd.Run()
	if nil != err {
		fmt.Println("err:", err)
	}
	fmt.Println(out.String())
}
