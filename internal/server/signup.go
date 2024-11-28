package server

import (
	"encoding/json"
	"net/http"
)

type SignupRequestBody struct {
	FirstName string `json:"first_name" validate:"required,alpha,ascii,min=2,max=25"`
	LastName  string `json:"last_name" validate:"required,alpha,ascii,min=2,max=25"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,ascii,min=8,max=50"`
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := validateAndGetValue[SignupRequestBody](r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Write data to response as json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
