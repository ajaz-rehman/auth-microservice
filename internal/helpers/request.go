package helpers

import (
	"context"
	"encoding/json"
	"io"

	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
)

func TransformAndValidateBody[T any](body io.ReadCloser) (value T, err error) {
	// Decode the json body into the value
	err = json.NewDecoder(body).Decode(&value)

	if err != nil {
		return
	}

	// Transform the value
	transform := modifiers.New()

	err = transform.Struct(context.Background(), &value)

	if err != nil {
		return
	}

	// Validate the value
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(&value)

	return
}
