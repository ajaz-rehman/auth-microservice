package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ajaz-rehman/auth-microservice/internal/controllers"
)

func controllerHandler(controller controllers.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		status, response, err := controller(r)

		if err != nil {
			errorHandler(w, status, err)
			return
		}

		err = responseHandler(w, status, response)

		if err != nil {
			errorHandler(w, http.StatusInternalServerError, err)
			return
		}
	}
}

func errorHandler(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errors := strings.Split(err.Error(), "\n")
	response := ErrorResponse{
		Errors: errors,
	}
	json.NewEncoder(w).Encode(response)
}

func responseHandler(w http.ResponseWriter, status int, response any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if response == nil {
		return nil
	}

	return json.NewEncoder(w).Encode(response)
}
