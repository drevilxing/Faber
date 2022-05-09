package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type PrometheusConfig struct {
	Global        Global          `yaml:"global"`
	ScrapeConfigs []ScrapeConfigs `yaml:"scrape_configs"`
}

type ExternalLabels struct {
	Monitor string `yaml:"monitor"`
}

type ScrapeConfigs struct {
	JobName       string          `yaml:"job_name"`
	StaticConfigs []StaticConfigs `yaml:"static_configs"`
}

type StaticConfigs struct {
	Targets []string `yaml:"targets,flow"`
}

type Global struct {
	ScrapeInterval string         `yaml:"scrape_interval"`
	ExternalLabels ExternalLabels `yaml:"external_labels"`
}

func ReadPrometheus() *PrometheusConfig {
	yamlFile, err := ioutil.ReadFile("./prometheus_template.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	var config *PrometheusConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(config)
	return config
}

func WritePrometheus(file *PrometheusConfig) {
	faber := ReadJsonFile()
	var nodes []string
	for i := 0; i < len(faber.Nodes); i++ {
		if faber.Nodes[i].Type[0] == "ca" {
			continue
		} else if faber.Nodes[i].Type[0] == "orderer" {
			nodes = append(nodes, faber.Nodes[i].Key+":"+"8443")
		} else {
			nodes = append(nodes, faber.Nodes[i].Key+":"+"9443")
		}
	}
	file.ScrapeConfigs[1].StaticConfigs[0].Targets = nodes
	data, err := yaml.Marshal(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	ioutil.WriteFile("./prometheus.yml", data, 0777)
	// fmt.Println(nodes)
}
