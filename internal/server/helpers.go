package server

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

func validateAndGetValue[T any](body io.ReadCloser) (value T, err error) {
	err = json.NewDecoder(body).Decode(&value)

	if err != nil {
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(value)

	return
}
