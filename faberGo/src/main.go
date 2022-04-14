package main

import "fmt"

func main() {
	environmentPath := "./environmentConfig.json"
	err := GenerateDeploymentConfigExample(environmentPath)
	if nil != err {
		fmt.Println(err.Error())
	}
	sdkPath := "./sdkConfig.json"
	err = GenerateGoSdkConfigExample(sdkPath)
	if nil != err {
		fmt.Println(err.Error())
	}
}
