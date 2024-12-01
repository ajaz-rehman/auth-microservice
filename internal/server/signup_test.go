package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestSignup(t *testing.T) {
	testUser := SignupRequestBody{
		FirstName: "test",
		LastName:  "user",
		Password:  "password",
		Email:     "test@gmail.com",
	}

	tests := []TableTest{
		{
			Name:           "Successful",
			Handler:        signupHandler,
			RequestPayload: testUser,
			ExpectedStatus: http.StatusCreated,
			ExpectedResponseFn: func(rr *httptest.ResponseRecorder) error {
				var tokens Tokens

				if err := json.NewDecoder(rr.Body).Decode(&tokens); err != nil {
					return err
				}

				if tokens.AccessToken == "" {
					return errors.New("empty access token")
				}

				if tokens.RefreshToken == "" {
					return errors.New("empty refresh token")
				}

				jwtSecret := os.Getenv("JWT_SECRET")

				userId, err := ValidateJWT(tokens.AccessToken, jwtSecret)

				if err != nil {
					return err
				}

				if userId != 1 {
					return errors.New("invalid user id")
				}

				return nil
			},
		},
		{
			Name:           "Duplicate",
			Handler:        signupHandler,
			RequestPayload: testUser,
			ExpectedStatus: http.StatusConflict,
		},
	}

	RunHttpTests(t, tests)
}
