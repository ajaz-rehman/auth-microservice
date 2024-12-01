package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TableTest struct {
	Name               string
	Handler            http.HandlerFunc
	RequestPayload     interface{}
	ExpectedStatus     int
	ExpectedResponseFn func(*httptest.ResponseRecorder) error
}

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

func RunHttpTests(t *testing.T, tests []TableTest) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			rr, err := MakeTestRequest(test.Handler, test.RequestPayload)

			if err != nil {
				t.Fatalf("Could not make request: %v", err)
			}

			if rr.Code != test.ExpectedStatus {
				t.Errorf("Expected status %v, got %v", test.ExpectedStatus, rr.Code)
			}

			if test.ExpectedResponseFn != nil {
				if err := test.ExpectedResponseFn(rr); err != nil {
					t.Errorf("Expected response failed: %v", err)
				}
			}
		})
	}
}
