package utils

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
