package core

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
