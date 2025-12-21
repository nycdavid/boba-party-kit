package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Search *Search `yaml:"search"`
	}

	Search struct {
		Data *Data `yaml:"data"`
	}

	Data struct {
		HTTP string `json:"http"`
		Auth *Auth  `json:"auth"`
	}

	Auth struct {
		Header *Header `json:"header"`
	}

	Header struct {
		Bearer string `json:"bearer"`
	}
)

func main() {
	b, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	var config Config
	if err := yaml.Unmarshal(b, &config); err != nil {
		panic(err)
	}

	if config.Search != nil {
		// Search component
		fmt.Println("Generating search component...")
	}
}
