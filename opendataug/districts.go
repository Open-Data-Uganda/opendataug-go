package opendataug

import (
	"fmt"
	"net/http"
)

// GetDistricts retrieves all districts
func (c *Client) GetDistricts() ([]District, error) {
	var response struct {
		Data []District `json:"data"`
	}

	err := c.doRequest(http.MethodGet, "/districts", &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetDistrict retrieves a specific district by ID
func (c *Client) GetDistrict(id string) (*District, error) {
	var response struct {
		Data District `json:"data"`
	}

	path := fmt.Sprintf("/districts/%s", id)
	err := c.doRequest(http.MethodGet, path, &response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}
