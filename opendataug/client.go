package opendataug

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var baseURL = "https://api.opendataug.com/v1"

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (c *Client) doRequest(method, path string, v interface{}) error {
	url := fmt.Sprintf("%s%s", baseURL, path)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return err
		}
	}

	return nil
}
