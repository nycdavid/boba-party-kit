package httpdriver

import (
	"io"
	"net/http"
	"os"
)

type (
	Client struct{}
)

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Get(url, authEnvVar string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	token := os.Getenv(authEnvVar)
	req.Header.Set("Authorization", "Bearer "+token)
	hc := &http.Client{}

	res, err := hc.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
