package core

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func MakeTestRequest(handle http.HandlerFunc, payload interface{}) (*httptest.ResponseRecorder, error) {
	body, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	// Method and URL are not important for testing since we are using a custom handler
	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	server := http.HandlerFunc(handle)
	rr := httptest.NewRecorder()
	server.ServeHTTP(rr, req)

	return rr, nil
}

func RunHttpTests(t *testing.T, tests []HttpTest) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			rr, err := MakeTestRequest(test.Handler, test.RequestPayload)

			if err != nil {
				t.Fatalf("Could not make request: %v", err)
			}

			if rr.Code != test.ExpectedStatus {
				t.Errorf("Expected status %v, got %v", test.ExpectedStatus, rr.Code)
			}

			if test.ExpectedResponse != nil {
				bodyBytes, err := io.ReadAll(rr.Body)

				if err != nil {
					t.Fatalf("Could not read response body: %v", err)
				}

				bodyString := strings.Trim(string(bodyBytes), "\n")

				if bodyString == "" {
					t.Fatalf("Empty response body")
				}

				expectedBytes, err := json.Marshal(test.ExpectedResponse)

				if err != nil {
					t.Fatalf("Could not marshal expected response: %v", err)
				}

				if bodyString != string(expectedBytes) {
					t.Errorf("Expected response %v, got %v", string(expectedBytes), bodyString)
				}
			}
		})
	}
}
