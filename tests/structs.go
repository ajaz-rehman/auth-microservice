package tests

import "net/http"

type TableTest struct {
	Name               string
	Method             string
	Endpoint           string
	RequestPayload     interface{}
	ExpectedStatus     int
	ExpectedResponseFn func(*http.Response) error
}
