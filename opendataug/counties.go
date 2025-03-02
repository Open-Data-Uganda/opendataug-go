package opendataug

import (
	"fmt"
	"net/http"
)

// GetCounties retrieves all counties
func (c *Client) GetCounties() ([]County, error) {
	var response struct {
		Data []County `json:"data"`
	}

	err := c.doRequest(http.MethodGet, "/counties", &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetCounty retrieves a specific county by ID
func (c *Client) GetCounty(id string) (*County, error) {
	var response struct {
		Data County `json:"data"`
	}

	path := fmt.Sprintf("/counties/%s", id)
	err := c.doRequest(http.MethodGet, path, &response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetCountiesByDistrict retrieves all counties in a specific district
func (c *Client) GetCountiesByDistrict(districtID string) ([]County, error) {
	var response struct {
		Data []County `json:"data"`
	}

	path := fmt.Sprintf("/districts/%s/counties", districtID)
	err := c.doRequest(http.MethodGet, path, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
