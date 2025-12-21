package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nycdavid/boba-party-kit/pkg/httpdriver"
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
		HTTP string `yaml:"http"`
		Auth *Auth  `yaml:"auth"`
	}

	Auth struct {
		Header *Header `yaml:"header"`
	}

	Header struct {
		BearerEnvVar string `yaml:"bearer-env-var"`
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
		hdc := httpdriver.NewClient()
		a, err := hdc.Get(
			config.Search.Data.HTTP,
			config.Search.Data.Auth.Header.BearerEnvVar,
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(a)

		p := tea.NewProgram(nil, tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
