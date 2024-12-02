package server

type ErrorResponse struct {
	Errors []string `json:"errors"`
}
