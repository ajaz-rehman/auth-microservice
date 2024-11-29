package core

import "net/http"

type ENV_TYPE string

const (
	DEVELOPMENT ENV_TYPE = "development"
	PRODUCTION  ENV_TYPE = "production"
	TEST        ENV_TYPE = "test"
)

type Env struct {
	PORT   string
	GO_ENV ENV_TYPE
}

type Config struct {
	Env Env
}

type HttpTest struct {
	Name             string
	Handler          http.HandlerFunc
	RequestPayload   interface{}
	ExpectedStatus   int
	ExpectedResponse interface{}
}
