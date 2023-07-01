package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	config := getConfig()
	for _, v := range config.Deployments {
		fmt.Println(v.Name + "," + v.Namespace)
		scale(&v)
	}
}
