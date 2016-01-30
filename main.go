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
	env, err := client.Info()
	if err != nil {
		log.Fatalln("Error getting docker info")
	}
	count := DockerCount{env.GetInt("Containers")}
	jsonCount, _ := json.Marshal(count)
	fmt.Println(string(jsonCount))
}
