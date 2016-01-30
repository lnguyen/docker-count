package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/fsouza/go-dockerclient"
)

type DockerCount struct {
	Containers int `json:"containers"`
}

func main() {
	endpoint := flag.String("endpoint", "unix:///var/run/docker.sock", "docker endpoint")
	flag.Parse()

	client, err := docker.NewClient(*endpoint)
	if err != nil {
		log.Fatalln("Error creating docker client")
	}
	containers, err := client.ListContainers(docker.ListContainersOptions{})
	if err != nil {
		log.Fatalln("Error getting docker info")
	}
	count := DockerCount{len(containers)}
	jsonCount, _ := json.Marshal(count)
	fmt.Println(string(jsonCount))
}
