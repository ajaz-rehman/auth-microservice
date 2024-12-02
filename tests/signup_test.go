package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ajaz-rehman/auth-microservice/internal/auth"
)

func TestSignupEndpoint(t *testing.T) {
	// Set up the server
	setupServer()

	// Define the test user
	testUser := map[string]string{
		"first_name": "test",
		"last_name":  "user",
		"email":      "test@gmail.com",
		"password":   "password",
	}

	// Marshal the test user to JSON
	body, err := json.Marshal(testUser)

	if err != nil {
		t.Fatalf("Could not marshal test user: %v", err)
	}

	// Make the request
	req, err := http.NewRequest("POST", "http://localhost:3000/signup", bytes.NewBuffer(body))

	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Record the response
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Could not make request: %v", err)
	}

	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status %v, got %v", http.StatusCreated, resp.StatusCode)
	}

	// Check the response body
	var tokens auth.Tokens
	err = json.NewDecoder(resp.Body).Decode(&tokens)

	if err != nil {
		t.Fatalf("Could not decode response: %v", err)
	}

	if tokens.AccessToken == "" {
		t.Errorf("Expected access token, got empty string")
	}

	if tokens.RefreshToken == "" {
		t.Errorf("Expected refresh token, got empty string")
	}
}
