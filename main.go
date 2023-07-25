package main

import (
	"log"
)

func main() {
	log.Println("start")

	config := getConfig()
	for _, v := range config.Deployments {
		scale(&v)
	}

	log.Println("end")
}
