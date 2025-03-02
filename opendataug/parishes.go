package opendataug

import (
	"fmt"
	"net/http"
)

// GetParishes retrieves all parishes
func (c *Client) GetParishes() ([]Parish, error) {
	var response struct {
		Data []Parish `json:"data"`
	}

	err := c.doRequest(http.MethodGet, "/parishes", &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetParish retrieves a specific parish by ID
func (c *Client) GetParish(id string) (*Parish, error) {
	var response struct {
		Data Parish `json:"data"`
	}

	path := fmt.Sprintf("/parishes/%s", id)
	err := c.doRequest(http.MethodGet, path, &response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetParishesBySubcounty retrieves all parishes in a specific subcounty
func (c *Client) GetParishesBySubcounty(subcountyID string) ([]Parish, error) {
	var response struct {
		Data []Parish `json:"data"`
	}

	path := fmt.Sprintf("/subcounties/%s/parishes", subcountyID)
	err := c.doRequest(http.MethodGet, path, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
