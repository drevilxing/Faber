package main

import "fmt"

func main() {
	path := "./config.json"
	err := GenerateDeploymentConfigExample(path)
	if nil != err {
		fmt.Println(err.Error())
	}
}
