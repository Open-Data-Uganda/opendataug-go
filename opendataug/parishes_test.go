package opendataug

import (
	"reflect"
	"testing"
)

func TestGetParishes(t *testing.T) {
	// Test response
	response := `{
		"data": [
			{
				"id": "parish-1",
				"name": "Kiwatule",
				"code": "KWT",
				"subcounty_id": "subcounty-1"
			},
			{
				"id": "parish-2",
				"name": "Bukoto",
				"code": "BKT",
				"subcounty_id": "subcounty-1"
			}
		]
	}`

	server, client := TestServer(t, "/parishes", response)
	defer server.Close()

	parishes, err := client.GetParishes()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := []Parish{
		{
			ID:          "parish-1",
			Name:        "Kiwatule",
			Code:        "KWT",
			SubcountyID: "subcounty-1",
		},
		{
			ID:          "parish-2",
			Name:        "Bukoto",
			Code:        "BKT",
			SubcountyID: "subcounty-1",
		},
	}

	if !reflect.DeepEqual(parishes, expected) {
		t.Errorf("Expected %+v, got %+v", expected, parishes)
	}
}

func TestGetParish(t *testing.T) {
	tests := []struct {
		name           string
		parishID       string
		expectedPath   string
		response       string
		expectedResult *Parish
		expectError    bool
	}{
		{
			name:         "Valid parish",
			parishID:     "parish-1",
			expectedPath: "/parishes/parish-1",
			response: `{
				"data": {
					"id": "parish-1",
					"name": "Kiwatule",
					"code": "KWT",
					"subcounty_id": "subcounty-1"
				}
			}`,
			expectedResult: &Parish{
				ID:          "parish-1",
				Name:        "Kiwatule",
				Code:        "KWT",
				SubcountyID: "subcounty-1",
			},
			expectError: false,
		},
		{
			name:         "Error response",
			parishID:     "invalid-id",
			expectedPath: "/parishes/invalid-id",
			response: `{
				"error": "Parish not found"
			}`,
			expectedResult: nil,
			expectError:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server, client := TestServer(t, tc.expectedPath, tc.response)
			defer server.Close()

			parish, err := client.GetParish(tc.parishID)

			if tc.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}

			if !tc.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			if !tc.expectError && !reflect.DeepEqual(parish, tc.expectedResult) {
				t.Errorf("Expected %+v, got %+v", tc.expectedResult, parish)
			}
		})
	}
}

func TestGetParishesBySubcounty(t *testing.T) {
	tests := []struct {
		name           string
		subcountyID    string
		expectedPath   string
		response       string
		expectedResult []Parish
		expectError    bool
	}{
		{
			name:         "Valid subcounty",
			subcountyID:  "subcounty-1",
			expectedPath: "/subcounties/subcounty-1/parishes",
			response: `{
				"data": [
					{
						"id": "parish-1",
						"name": "Kiwatule",
						"code": "KWT",
						"subcounty_id": "subcounty-1"
					},
					{
						"id": "parish-2",
						"name": "Bukoto",
						"code": "BKT",
						"subcounty_id": "subcounty-1"
					}
				]
			}`,
			expectedResult: []Parish{
				{
					ID:          "parish-1",
					Name:        "Kiwatule",
					Code:        "KWT",
					SubcountyID: "subcounty-1",
				},
				{
					ID:          "parish-2",
					Name:        "Bukoto",
					Code:        "BKT",
					SubcountyID: "subcounty-1",
				},
			},
			expectError: false,
		},
		{
			name:         "Empty result",
			subcountyID:  "subcounty-2",
			expectedPath: "/subcounties/subcounty-2/parishes",
			response: `{
				"data": []
			}`,
			expectedResult: []Parish{},
			expectError:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server, client := TestServer(t, tc.expectedPath, tc.response)
			defer server.Close()

			parishes, err := client.GetParishesBySubcounty(tc.subcountyID)

			if tc.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}

			if !tc.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			if !tc.expectError && !reflect.DeepEqual(parishes, tc.expectedResult) {
				t.Errorf("Expected %+v, got %+v", tc.expectedResult, parishes)
			}
		})
	}
}
