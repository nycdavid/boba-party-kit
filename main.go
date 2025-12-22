package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nycdavid/boba-party-kit/internal/config"
	"github.com/nycdavid/boba-party-kit/internal/views/layout"
	"gopkg.in/yaml.v3"
)

func main() {
	b, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	cfg := &config.Config{}
	if err := yaml.Unmarshal(b, cfg); err != nil {
		panic(err)
	}

	if cfg.Search != nil {
		// Search component
		//hdc := httpdriver.NewClient()
		//a, err := hdc.Get(
		//	config.Search.Init.Data.HTTP,
		//	config.Search.Init.Data.Auth.Header.BearerEnvVar,
		//)
		//if err != nil {
		//	log.Fatal(err)
		//}

		l := layout.New(cfg)
		p := tea.NewProgram(l, tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
