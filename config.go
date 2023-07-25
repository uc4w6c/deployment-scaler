package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Deployment struct {
	Namespace string
	Name      string
	Replicas  int
}

type Config struct {
	Deployments []Deployment
}

func getConfig() Config {
	config := Config{}
	b, _ := os.ReadFile("config.yaml")

	err := yaml.Unmarshal(b, &config)
	if err != nil {
		log.Panic(err)
	}
	return config
}
