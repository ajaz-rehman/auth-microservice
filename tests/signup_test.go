package tests

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/ajaz-rehman/auth-microservice/internal/auth"
	"github.com/ajaz-rehman/auth-microservice/internal/handlers"
)

func TestSignup(t *testing.T) {
	setupServer()

	testUser := handlers.SignupRequest{
		FirstName: "test",
		LastName:  "user",
		Password:  "password",
		Email:     "test@gmail.com",
	}

	tests := []TableTest{
		{
			Name:           "Successful",
			Method:         "POST",
			Endpoint:       "/signup",
			RequestPayload: testUser,
			ExpectedStatus: http.StatusCreated,
			ExpectedResponseFn: func(resp *http.Response) error {
				var tokens auth.Tokens

				if err := json.NewDecoder(resp.Body).Decode(&tokens); err != nil {
					return err
				}

				if tokens.AccessToken == "" {
					return errors.New("empty access token")
				}

				if tokens.RefreshToken == "" {
					return errors.New("empty refresh token")
				}

				jwtSecret := os.Getenv("JWT_SECRET")

				userId, err := auth.ValidateJWTToken(tokens.AccessToken, jwtSecret)

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
			Method:         "POST",
			Endpoint:       "/signup",
			RequestPayload: testUser,
			ExpectedStatus: http.StatusConflict,
		},
	}

	RunHttpTests(t, tests)
}
