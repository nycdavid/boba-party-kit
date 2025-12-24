package datadriver

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/nycdavid/boba-party-kit/internal/config"
)

type (
	HTTP struct {
		cfg config.HTTP

		url        string
		authEnvVar string
		method     string
	}
)

func NewHTTP(url, authEnvVar, method string) *HTTP {
	return &HTTP{
		url:        url,
		authEnvVar: authEnvVar,
		method:     method,
	}
}

func (h *HTTP) Fetch() ([]byte, error) {
	req, err := http.NewRequest(h.method, h.url, nil)
	if err != nil {
		return nil, err
	}

	token := os.Getenv(h.authEnvVar)
	req.Header.Set("Authorization", "Bearer "+token)
	hc := &http.Client{}

	res, err := hc.Do(req)
	defer func() {
		if err := res.Body.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
