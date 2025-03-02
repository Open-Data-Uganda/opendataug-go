package opendataug

import (
	"fmt"
	"net/http"
)

// GetVillages retrieves all villages
func (c *Client) GetVillages() ([]Village, error) {
	var response struct {
		Data []Village `json:"data"`
	}

	err := c.doRequest(http.MethodGet, "/villages", &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetVillage retrieves a specific village by ID
func (c *Client) GetVillage(id string) (*Village, error) {
	var response struct {
		Data Village `json:"data"`
	}

	path := fmt.Sprintf("/villages/%s", id)
	err := c.doRequest(http.MethodGet, path, &response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetVillagesByParish retrieves all villages in a specific parish
func (c *Client) GetVillagesByParish(parishID string) ([]Village, error) {
	var response struct {
		Data []Village `json:"data"`
	}

	path := fmt.Sprintf("/parishes/%s/villages", parishID)
	err := c.doRequest(http.MethodGet, path, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
