package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func RunHttpTests(t *testing.T, tests []TableTest) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			resp, err := MakeRequest(test.Method, test.Endpoint, test.RequestPayload)

			if err != nil {
				t.Fatalf("Could not make request: %v", err)
			}

			defer resp.Body.Close()

			if resp.StatusCode != test.ExpectedStatus {
				t.Errorf("Expected status %v, got %v", test.ExpectedStatus, resp.StatusCode)
			} else if test.ExpectedResponseFn != nil {
				if err := test.ExpectedResponseFn(resp); err != nil {
					t.Errorf("Expected response failed: %v", err)
				}
			}
		})
	}
}

func MakeRequest(method string, endpoint string, payload interface{}) (*http.Response, error) {
	body, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	url := getServerURL() + endpoint

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
