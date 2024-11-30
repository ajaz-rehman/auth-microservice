package server

import (
	"net/http"
	"testing"

	"github.com/ajaz-rehman/auth-microservice/internal/core"
)

func TestSignup(t *testing.T) {
	testUser := SignupRequestBody{
		FirstName: "test",
		LastName:  "user",
		Password:  "password",
		Email:     "test@gmail.com",
	}

	tokens := Tokens{
		AccessToken:  "access",
		RefreshToken: "refresh",
	}

	tests := []core.HttpTest{
		{
			Name:             "Successful",
			Handler:          signupHandler,
			RequestPayload:   testUser,
			ExpectedStatus:   http.StatusCreated,
			ExpectedResponse: tokens,
		},
		{
			Name:             "Duplicate",
			Handler:          signupHandler,
			RequestPayload:   testUser,
			ExpectedStatus:   http.StatusConflict,
			ExpectedResponse: nil,
		},
	}

	core.RunHttpTests(t, tests)
}
