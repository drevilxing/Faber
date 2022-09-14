package main

import "faber/pkg/configtx"

func main() {
	println("Start!")
	configtx.GenerateConfigTx().Export("./")
	println("End!")
}
