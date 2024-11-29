package server

import (
	"net/http"
	"testing"

	"github.com/ajaz-rehman/auth-microservice/internal/core"
)

func TestSignup(t *testing.T) {
	tests := []core.HttpTest{
		{
			Name:    "Successful",
			Handler: signupHandler,
			RequestPayload: SignupRequestBody{
				FirstName: "test",
				LastName:  "user",
				Password:  "password",
				Email:     "test@gmail.com",
			},
			ExpectedStatus:   http.StatusCreated,
			ExpectedResponse: nil,
		},
		{
			Name:    "Duplicate",
			Handler: signupHandler,
			RequestPayload: SignupRequestBody{
				FirstName: "test",
				LastName:  "user",
				Password:  "password",
				Email:     "test@gmail.com",
			},
			ExpectedStatus:   http.StatusConflict,
			ExpectedResponse: nil,
		},
	}

	core.RunHttpTests(t, tests)
}
