1 Go Backend https Server + Python -> yaml
    - p1
    - p2
    - p1 SDK (Config, binary)
2 Project1 /root/opt
3 Project2 /root/opt


docker command
```Shell
docker run -itd --name projects -v /var/run:/var/run -v /Users/mac/Documents/Projects:/projects debian
docker exec -it projects /bin/bash

```

In docker container
```Shell
apt-get update -y
apt-get install apt-transport-https ca-certificates -y
mv /etc/apt/sources.list /etc/apt/sources.list.bak
echo "deb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye main contrib non-free
# deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye main contrib non-free
deb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-updates main contrib non-free
# deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-updates main contrib non-free

deb https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-backports main contrib non-free
# deb-src https://mirrors.tuna.tsinghua.edu.cn/debian/ bullseye-backports main contrib non-free

deb https://mirrors.tuna.tsinghua.edu.cn/debian-security bullseye-security main contrib non-free
# deb-src https://mirrors.tuna.tsinghua.edu.cn/debian-security bullseye-security main contrib non-free" >> /etc/apt/sources.list

apt-get update -y

apt-get install vim wget git docker docker.io docker-compose -y

wget https://studygolang.com/dl/golang/go1.17.3.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.3.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
source ~/.profile
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

docker pull hyperledger/fabric-peer:2.2.0
docker pull hyperledger/fabric-orderer:2.2.0
docker pull hyperledger/fabric-ca:1.4.7
#docker pull hyperledger/fabric-tools:2.2.0
#docker pull hyperledger/fabric-ccenv:2.2.0
#docker pull hyperledger/fabric-baseos:2.2.0
#docker pull hyperledger/fabric-javaenv:2.2.0
#docker pull hyperledger/fabric-nodeenv:2.2.0
```

```Shell
git clone https://github.com/EternallyAscend/Faber
cd Faber/faberGo


```

Tsinghua Debian

```Shell
docker stop projects
docker rm projects
docker volume prune
```
