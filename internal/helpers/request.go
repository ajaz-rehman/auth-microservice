package helpers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
)

func HandleRequest[T any](fn RequestHandlerFn[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		data, err := transformAndValidateBody[T](r.Body)

		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		status, resp, err := fn(data, r)

		if err != nil {
			if strings.Contains(err.Error(), "violates unique constraint") {
				status = http.StatusConflict
			} else if status == 0 {
				status = http.StatusInternalServerError
			}

			respondWithError(w, status, err)
			return
		}

		if status == 0 {
			status = http.StatusOK
		}

		respondWithJSON(w, status, resp)
	}
}

func transformAndValidateBody[T any](body io.ReadCloser) (value T, err error) {
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
