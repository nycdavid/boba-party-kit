package config

import "fmt"

type (
	HTTP struct {
		URL    string `yaml:"url"`
		Method string `yaml:"method"` // defaults to GET
		Auth   *Auth  `yaml:"auth"`
	}
)

func (h *HTTP) FormattedURL(parts []any) string {
	return fmt.Sprintf(h.URL, parts...)
}
