package opendataug

import (
	"reflect"
	"testing"
)

func TestGetVillages(t *testing.T) {
	// Test response
	response := `{
		"data": [
			{
				"id": "village-1",
				"name": "Kiwatule Central",
				"code": "KWTC",
				"parish_id": "parish-1"
			},
			{
				"id": "village-2",
				"name": "Kiwatule East",
				"code": "KWTE",
				"parish_id": "parish-1"
			}
		]
	}`

	server, client := TestServer(t, "/villages", response)
	defer server.Close()

	villages, err := client.GetVillages()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := []Village{
		{
			ID:       "village-1",
			Name:     "Kiwatule Central",
			Code:     "KWTC",
			ParishID: "parish-1",
		},
		{
			ID:       "village-2",
			Name:     "Kiwatule East",
			Code:     "KWTE",
			ParishID: "parish-1",
		},
	}

	if !reflect.DeepEqual(villages, expected) {
		t.Errorf("Expected %+v, got %+v", expected, villages)
	}
}

func TestGetVillage(t *testing.T) {
	tests := []struct {
		name           string
		villageID      string
		expectedPath   string
		response       string
		expectedResult *Village
		expectError    bool
	}{
		{
			name:         "Valid village",
			villageID:    "village-1",
			expectedPath: "/villages/village-1",
			response: `{
				"data": {
					"id": "village-1",
					"name": "Kiwatule Central",
					"code": "KWTC",
					"parish_id": "parish-1"
				}
			}`,
			expectedResult: &Village{
				ID:       "village-1",
				Name:     "Kiwatule Central",
				Code:     "KWTC",
				ParishID: "parish-1",
			},
			expectError: false,
		},
		{
			name:         "Error response",
			villageID:    "invalid-id",
			expectedPath: "/villages/invalid-id",
			response: `{
				"error": "Village not found"
			}`,
			expectedResult: nil,
			expectError:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server, client := TestServer(t, tc.expectedPath, tc.response)
			defer server.Close()

			village, err := client.GetVillage(tc.villageID)

			if tc.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}

			if !tc.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			if !tc.expectError && !reflect.DeepEqual(village, tc.expectedResult) {
				t.Errorf("Expected %+v, got %+v", tc.expectedResult, village)
			}
		})
	}
}

func TestGetVillagesByParish(t *testing.T) {
	tests := []struct {
		name           string
		parishID       string
		expectedPath   string
		response       string
		expectedResult []Village
		expectError    bool
	}{
		{
			name:         "Valid parish",
			parishID:     "parish-1",
			expectedPath: "/parishes/parish-1/villages",
			response: `{
				"data": [
					{
						"id": "village-1",
						"name": "Kiwatule Central",
						"code": "KWTC",
						"parish_id": "parish-1"
					},
					{
						"id": "village-2",
						"name": "Kiwatule East",
						"code": "KWTE",
						"parish_id": "parish-1"
					}
				]
			}`,
			expectedResult: []Village{
				{
					ID:       "village-1",
					Name:     "Kiwatule Central",
					Code:     "KWTC",
					ParishID: "parish-1",
				},
				{
					ID:       "village-2",
					Name:     "Kiwatule East",
					Code:     "KWTE",
					ParishID: "parish-1",
				},
			},
			expectError: false,
		},
		{
			name:         "Empty result",
			parishID:     "parish-2",
			expectedPath: "/parishes/parish-2/villages",
			response: `{
				"data": []
			}`,
			expectedResult: []Village{},
			expectError:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			server, client := TestServer(t, tc.expectedPath, tc.response)
			defer server.Close()

			villages, err := client.GetVillagesByParish(tc.parishID)

			if tc.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}

			if !tc.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			if !tc.expectError && !reflect.DeepEqual(villages, tc.expectedResult) {
				t.Errorf("Expected %+v, got %+v", tc.expectedResult, villages)
			}
		})
	}
}
