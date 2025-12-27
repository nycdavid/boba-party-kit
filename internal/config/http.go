package config

import "fmt"

type (
	HTTP struct {
		URL    string `yaml:"url"`
		Method string `yaml:"method"` // defaults to GET
		Auth   *Auth  `yaml:"auth"`
	}
)

func (h *HTTP) FormattedURL(parts []string) string {
	anys := make([]any, len(parts))
	for i, p := range parts {
		anys[i] = p
	}
	return fmt.Sprintf(h.URL, anys...)
}
