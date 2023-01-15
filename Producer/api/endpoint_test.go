package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Test ping function
func TestPing(t *testing.T) {
	// Create a new router
	router := gin.Default()
	// Register the route
	router.GET("/ping", Ping)
	// Create a new request
	req, _ := http.NewRequest("GET", "/ping", nil)
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

func TestHealth(t *testing.T) {
	// Create a new router
	router := gin.Default()
	// Register the route
	router.GET("/health", Health)
	// Create a new request
	req, _ := http.NewRequest("GET", "/health", nil)
	// Create a new response recorder
	res := httptest.NewRecorder()
	// Perform the request
	router.ServeHTTP(res, req)
	// Check the response status code is what we expect
	if res.Code != 200 {
		t.Errorf("Expected status code 200, got %d", res.Code)
	}
	// Check the response body is what we expect
	expected := `{"status":"UP","message":"Service is running"}`
	if res.Body.String() != expected {
		t.Errorf("Expected response body to be %s, got %s", expected, res.Body.String())
	}
}
