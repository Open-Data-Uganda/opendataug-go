package opendataug

import (
	"fmt"
	"net/http"
)

// GetSubcounties retrieves all subcounties
func (c *Client) GetSubcounties() ([]Subcounty, error) {
	var response struct {
		Data []Subcounty `json:"data"`
	}

	err := c.doRequest(http.MethodGet, "/subcounties", &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetSubcounty retrieves a specific subcounty by ID
func (c *Client) GetSubcounty(id string) (*Subcounty, error) {
	var response struct {
		Data Subcounty `json:"data"`
	}

	path := fmt.Sprintf("/subcounties/%s", id)
	err := c.doRequest(http.MethodGet, path, &response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetSubcountiesByCounty retrieves all subcounties in a specific county
func (c *Client) GetSubcountiesByCounty(countyID string) ([]Subcounty, error) {
	var response struct {
		Data []Subcounty `json:"data"`
	}

	path := fmt.Sprintf("/counties/%s/subcounties", countyID)
	err := c.doRequest(http.MethodGet, path, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
