package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test Gin router
func TestApiRouter(t *testing.T) {
	// Create a new router
	router := ApiV1()
	// Create a new request
	req, _ := http.NewRequest("GET", "/api/v1/ping", nil)
	// Create a new response recorder
	res := httptest.NewRecorder()
	// Perform the request
	router.ServeHTTP(res, req)
	// Check the response status code is what we expect
	if res.Code != 200 {
		t.Errorf("Expected status code 200, got %d", res.Code)
	}
	// Check the response body is what we expect
	expected := `{"message":"pong"}`
	if res.Body.String() != expected {
		t.Errorf("Expected response body to be %s, got %s", expected, res.Body.String())
	}
}
