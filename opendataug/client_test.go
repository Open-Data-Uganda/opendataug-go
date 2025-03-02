package opendataug

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	apiKey := "test-api-key"
	client := NewClient(apiKey)

	if client.apiKey != apiKey {
		t.Errorf("Expected apiKey to be %s, got %s", apiKey, client.apiKey)
	}

	if client.httpClient == nil {
		t.Error("Expected httpClient to be initialized")
	}

	if client.httpClient.Timeout != time.Second*30 {
		t.Errorf("Expected timeout to be 30 seconds, got %v", client.httpClient.Timeout)
	}
}

func TestDoRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("x-api-key") != "test-api-key" {
			t.Errorf("Expected x-api-key header to be test-api-key, got %s", r.Header.Get("x-api-key"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type header to be application/json, got %s", r.Header.Get("Content-Type"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"test": "data"}`))
	}))
	defer server.Close()

	originalBaseURL := baseURL
	baseURL = server.URL
	defer func() { baseURL = originalBaseURL }()

	client := NewClient("test-api-key")

	var response map[string]string
	err := client.doRequest(http.MethodGet, "/test", &response)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if response["test"] != "data" {
		t.Errorf("Expected response to be {test: data}, got %v", response)
	}
}

func TestDoRequestError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	originalBaseURL := baseURL
	baseURL = server.URL
	defer func() { baseURL = originalBaseURL }()

	client := NewClient("test-api-key")

	var response map[string]string
	err := client.doRequest(http.MethodGet, "/test", &response)
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

func TestDoRequestInvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"test": "data"`))
	}))
	defer server.Close()

	originalBaseURL := baseURL
	baseURL = server.URL
	defer func() { baseURL = originalBaseURL }()

	client := NewClient("test-api-key")

	var response map[string]string
	err := client.doRequest(http.MethodGet, "/test", &response)
	if err == nil {
		t.Error("Expected an error for invalid JSON, got nil")
	}
}
