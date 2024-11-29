package server

import (
	"net/http"
	"testing"

	"github.com/ajaz-rehman/auth-microservice/internal/core"
)

func TestSignup(t *testing.T) {
	payload := SignupRequestBody{
		FirstName: "test",
		LastName:  "user",
		Password:  "password",
		Email:     "test@gmail.com",
	}

	rr, err := core.MakeTestRequest(signupHandler, payload)

	if err != nil {
		t.Fatalf("Could not make request: %v", err)
	}

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}
