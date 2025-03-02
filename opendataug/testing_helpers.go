package opendataug

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T, expectedPath string, response string) (*httptest.Server, *Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != expectedPath {
			t.Errorf("Expected request to %s, got %s", expectedPath, r.URL.Path)
		}

		if r.Header.Get("x-api-key") != "test-api-key" {
			t.Errorf("Expected API key header to be set")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))

	client := NewClient("test-api-key")

	baseURL = server.URL

	return server, client
}
