package opendataug

import (
	"reflect"
	"testing"
)

func TestGetSubcounties(t *testing.T) {
	// Test response
	response := `{
		"data": [
			{
				"id": "subcounty-1",
				"name": "Ntinda",
				"code": "NTD",
				"county_id": "county-1"
			},
			{
				"id": "subcounty-2",
				"name": "Kyambogo",
				"code": "KYB",
				"county_id": "county-1"
			}
		]
	}`

	server, client := TestServer(t, "/subcounties", response)
	defer server.Close()

	subcounties, err := client.GetSubcounties()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := []Subcounty{
		{
			ID:       "subcounty-1",
			Name:     "Ntinda",
			Code:     "NTD",
			CountyID: "county-1",
		},
		{
			ID:       "subcounty-2",
			Name:     "Kyambogo",
			Code:     "KYB",
			CountyID: "county-1",
		},
	}

	if !reflect.DeepEqual(subcounties, expected) {
		t.Errorf("Expected %+v, got %+v", expected, subcounties)
	}
}

func TestGetSubcounty(t *testing.T) {
	tests := []struct {
		name           string
		subcountyID    string
		expectedPath   string
		response       string
		expectedResult *Subcounty
		expectError    bool
	}{
		{
			name:         "Valid subcounty",
			subcountyID:  "subcounty-1",
			expectedPath: "/subcounties/subcounty-1",
			response: `{
				"data": {
					"id": "subcounty-1",
					"name": "Ntinda",
					"code": "NTD",
					"county_id": "county-1"
				}
			}`,
			expectedResult: &Subcounty{
				ID:       "subcounty-1",
				Name:     "Ntinda",
				Code:     "NTD",
				CountyID: "county-1",
			},
			expectError: false,
		},
		{
			name:         "Error response",
			subcountyID:  "invalid-id",
			expectedPath: "/subcounties/invalid-id",
			response: `{
				"error": "Subcounty not found"
			}`,
			expectedResult: nil,
			expectError:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server, client := TestServer(t, tc.expectedPath, tc.response)
			defer server.Close()

			subcounty, err := client.GetSubcounty(tc.subcountyID)

			if tc.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}

			if !tc.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			if !tc.expectError && !reflect.DeepEqual(subcounty, tc.expectedResult) {
				t.Errorf("Expected %+v, got %+v", tc.expectedResult, subcounty)
			}
		})
	}
}

func TestGetSubcountiesByCounty(t *testing.T) {
	tests := []struct {
		name           string
		countyID       string
		expectedPath   string
		response       string
		expectedResult []Subcounty
		expectError    bool
	}{
		{
			name:         "Valid county",
			countyID:     "county-1",
			expectedPath: "/counties/county-1/subcounties",
			response: `{
				"data": [
					{
						"id": "subcounty-1",
						"name": "Ntinda",
						"code": "NTD",
						"county_id": "county-1"
					},
					{
						"id": "subcounty-2",
						"name": "Kyambogo",
						"code": "KYB",
						"county_id": "county-1"
					}
				]
			}`,
			expectedResult: []Subcounty{
				{
					ID:       "subcounty-1",
					Name:     "Ntinda",
					Code:     "NTD",
					CountyID: "county-1",
				},
				{
					ID:       "subcounty-2",
					Name:     "Kyambogo",
					Code:     "KYB",
					CountyID: "county-1",
				},
			},
			expectError: false,
		},
		{
			name:         "Empty result",
			countyID:     "county-2",
			expectedPath: "/counties/county-2/subcounties",
			response: `{
				"data": []
			}`,
			expectedResult: []Subcounty{},
			expectError:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server, client := TestServer(t, tc.expectedPath, tc.response)
			defer server.Close()

			subcounties, err := client.GetSubcountiesByCounty(tc.countyID)

			if tc.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}

			if !tc.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			if !tc.expectError && !reflect.DeepEqual(subcounties, tc.expectedResult) {
				t.Errorf("Expected %+v, got %+v", tc.expectedResult, subcounties)
			}
		})
	}
}
