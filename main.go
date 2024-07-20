package main

import (
	"gopkg.in/yaml.v3"
	"http-server-go/controller"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	App struct {
		Name string `yaml:"name"`
	} `yaml:"app"`
}

// private function
func readConfig() (*Config, error) {
	var config Config
	data, err := os.ReadFile("config.yaml")

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return &config, nil
}

func main() {
	config, _ := readConfig()
	log.Printf("Server started on port %s", config.Server.Port)

	http.HandleFunc("/randomElement", controller.RandomElement)
	http.HandleFunc("/element/{id}", controller.IndexElement)

	err := http.ListenAndServe(":"+config.Server.Port, nil)
	if err != nil {
		panic(err)
	}

}
