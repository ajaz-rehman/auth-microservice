package server

import (
	"encoding/json"
	"net/http"
)

type SignupRequestBody struct {
	FirstName string `json:"first_name" mod:"trim,lcase,title" validate:"required,alpha,ascii,min=2,max=25"`
	LastName  string `json:"last_name" mod:"trim,lcase,title" validate:"required,alpha,ascii,min=2,max=25"`
	Email     string `json:"email" mod:"trim,lcase" validate:"required,email"`
	Password  string `json:"password" validate:"required,ascii,min=8,max=50,excludes= "`
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := transformAndValidateBody[SignupRequestBody](r.Body)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
