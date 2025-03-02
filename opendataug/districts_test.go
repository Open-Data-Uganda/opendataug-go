package opendataug

import (
	"reflect"
	"testing"
)

func TestGetDistricts(t *testing.T) {
	// Test response
	response := `{
		"data": [
			{
				"id": "district-1",
				"name": "Kampala",
				"town_status": true,
				"region_id": "region-1",
				"region_name": "Central"
			},
			{
				"id": "district-2",
				"name": "Wakiso",
				"town_status": false,
				"region_id": "region-1",
				"region_name": "Central"
			}
		]
	}`

	server, client := TestServer(t, "/districts", response)
	defer server.Close()

	districts, err := client.GetDistricts()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := []District{
		{
			ID:         "district-1",
			Name:       "Kampala",
			TownStatus: true,
			RegionID:   "region-1",
			RegionName: "Central",
		},
		{
			ID:         "district-2",
			Name:       "Wakiso",
			TownStatus: false,
			RegionID:   "region-1",
			RegionName: "Central",
		},
	}

	if !reflect.DeepEqual(districts, expected) {
		t.Errorf("Expected %+v, got %+v", expected, districts)
	}
}

func TestGetDistrict(t *testing.T) {
	tests := []struct {
		name           string
		districtID     string
		expectedPath   string
		response       string
		expectedResult *District
		expectError    bool
	}{
		{
			name:         "Valid district",
			districtID:   "district-1",
			expectedPath: "/districts/district-1",
			response: `{
				"data": {
					"id": "district-1",
					"name": "Kampala",
					"town_status": true,
					"region_id": "region-1",
					"region_name": "Central"
				}
			}`,
			expectedResult: &District{
				ID:         "district-1",
				Name:       "Kampala",
				TownStatus: true,
				RegionID:   "region-1",
				RegionName: "Central",
			},
			expectError: false,
		},
		{
			name:         "Error response",
			districtID:   "invalid-id",
			expectedPath: "/districts/invalid-id",
			response: `{
				"error": "District not found"
			}`,
			expectedResult: nil,
			expectError:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server, client := TestServer(t, tc.expectedPath, tc.response)
			defer server.Close()

			district, err := client.GetDistrict(tc.districtID)

			if tc.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}

			if !tc.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			if !tc.expectError && !reflect.DeepEqual(district, tc.expectedResult) {
				t.Errorf("Expected %+v, got %+v", tc.expectedResult, district)
			}
		})
	}
}
