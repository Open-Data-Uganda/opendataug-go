package opendataug

import (
	"reflect"
	"testing"
)

func TestGetCounties(t *testing.T) {
	response := `{
		"data": [
			{
				"id": "county-1",
				"name": "Nakawa",
				"code": "NKW",
				"district_id": "district-1"
			},
			{
				"id": "county-2",
				"name": "Kawempe",
				"code": "KWP",
				"district_id": "district-1"
			}
		]
	}`

	server, client := TestServer(t, "/counties", response)
	defer server.Close()

	counties, err := client.GetCounties()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := []County{
		{
			ID:         "county-1",
			Name:       "Nakawa",
			Code:       "NKW",
			DistrictID: "district-1",
		},
		{
			ID:         "county-2",
			Name:       "Kawempe",
			Code:       "KWP",
			DistrictID: "district-1",
		},
	}

	if !reflect.DeepEqual(counties, expected) {
		t.Errorf("Expected %+v, got %+v", expected, counties)
	}
}

func TestGetCounty(t *testing.T) {
	tests := []struct {
		name           string
		countyID       string
		expectedPath   string
		response       string
		expectedResult *County
		expectError    bool
	}{
		{
			name:         "Valid county",
			countyID:     "county-1",
			expectedPath: "/counties/county-1",
			response: `{
				"data": {
					"id": "county-1",
					"name": "Nakawa",
					"code": "NKW",
					"district_id": "district-1"
				}
			}`,
			expectedResult: &County{
				ID:         "county-1",
				Name:       "Nakawa",
				Code:       "NKW",
				DistrictID: "district-1",
			},
			expectError: false,
		},
		{
			name:         "Error response",
			countyID:     "invalid-id",
			expectedPath: "/counties/invalid-id",
			response: `{
				"error": "County not found"
			}`,
			expectedResult: nil,
			expectError:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server, client := TestServer(t, tc.expectedPath, tc.response)
			defer server.Close()

			county, err := client.GetCounty(tc.countyID)

			if tc.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}

			if !tc.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			if !tc.expectError && !reflect.DeepEqual(county, tc.expectedResult) {
				t.Errorf("Expected %+v, got %+v", tc.expectedResult, county)
			}
		})
	}
}

func TestGetCountiesByDistrict(t *testing.T) {
	tests := []struct {
		name           string
		districtID     string
		expectedPath   string
		response       string
		expectedResult []County
		expectError    bool
	}{
		{
			name:         "Valid district",
			districtID:   "district-1",
			expectedPath: "/districts/district-1/counties",
			response: `{
				"data": [
					{
						"id": "county-1",
						"name": "Nakawa",
						"code": "NKW",
						"district_id": "district-1"
					},
					{
						"id": "county-2",
						"name": "Kawempe",
						"code": "KWP",
						"district_id": "district-1"
					}
				]
			}`,
			expectedResult: []County{
				{
					ID:         "county-1",
					Name:       "Nakawa",
					Code:       "NKW",
					DistrictID: "district-1",
				},
				{
					ID:         "county-2",
					Name:       "Kawempe",
					Code:       "KWP",
					DistrictID: "district-1",
				},
			},
			expectError: false,
		},
		{
			name:         "Empty result",
			districtID:   "district-2",
			expectedPath: "/districts/district-2/counties",
			response: `{
				"data": []
			}`,
			expectedResult: []County{},
			expectError:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server, client := TestServer(t, tc.expectedPath, tc.response)
			defer server.Close()

			counties, err := client.GetCountiesByDistrict(tc.districtID)

			if tc.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}

			if !tc.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			if !tc.expectError && !reflect.DeepEqual(counties, tc.expectedResult) {
				t.Errorf("Expected %+v, got %+v", tc.expectedResult, counties)
			}
		})
	}
}
