package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type Faber struct {
	Key         string        `json:"key"`
	Groups      []Groups      `json:"groups"`
	Nodes       []Nodes       `json:"nodes"`
	Blockchains []Blockchains `json:"blockchains"`
}
type Node struct {
	Ca              string        `json:"ca"`
	Orderer         []string      `json:"orderer"`
	LeaderPeers     []interface{} `json:"leader_peers"`
	AnchorPeers     []interface{} `json:"anchor_peers"`
	CommittingPeers []interface{} `json:"committing_peers"`
	EndorsingPeers  []interface{} `json:"endorsing_peers"`
}
type Groups struct {
	Key         string   `json:"key"`
	Ca          string   `json:"ca"`
	Node        Node     `json:"node"`
	Blockchains string   `json:"blockchains"`
	Channel     []string `json:"channel"`
}
type Address struct {
	Host       string `json:"host"`
	SSHPort    string `json:"ssh_port"`
	FabricPort string `json:"fabric_port"`
}
type Nodes struct {
	Key       string        `json:"key"`
	Org       string        `json:"org"`
	Address   Address       `json:"address"`
	Bootstrap []interface{} `json:"bootstrap"`
	Type      []string      `json:"type"`
}
type Blockchains struct {
	Key      string   `json:"key"`
	Name     string   `json:"name"`
	Channels []string `json:"channels"`
}

func getstatus(host string, port string) bool {
	cmd := exec.Command("/bin/bash", "-c", "curl"+host+":"+port+"/"+"healthz")

	output, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("无法获取命令的标准输出管道", err.Error())
		return false
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Linux命令执行失败，请检查命令输入是否有误", err.Error())
		return false
	}

	bytes, err := ioutil.ReadAll(output)
	if err != nil {
		fmt.Println("打印异常，请检查")
		return false
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait", err.Error())
		return false
	}

	//fmt.Printf("打印内存信息：\n\n%s", bytes)
	return isOK(bytes)
}

func isOK(status []byte) bool {
	var v interface{}
	json.Unmarshal(status, &v)
	data := v.(map[string]interface{})

	fmt.Println(data["status"])
	return data["status"] == "OK"
}

func ReadJsonFile() *Faber {
	jsonFile, err := os.Open("faber.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened faber.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var faber Faber

	json.Unmarshal(byteValue, &faber)
	return &faber
}

func main() {
	file := ReadPrometheus()
	WritePrometheus(file)

	workDir, _ := os.Getwd()

	RunContainer("prom/prometheus", "/prometheus", "9090", "9090", map[string]string{workDir + "/prometheus.yml": "/etc/prometheus/prometheus.yml"})
	ConnectToNecwork("/prometheus")

	RunContainer("grafana/grafana", "/grafana", "3000", "3000", map[string]string{})
	ConnectToNecwork("/grafana")

}
