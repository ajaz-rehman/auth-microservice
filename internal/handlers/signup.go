package handlers

import (
	"net/http"

	"github.com/ajaz-rehman/auth-microservice/internal/app"
	"github.com/ajaz-rehman/auth-microservice/internal/auth"
	"github.com/ajaz-rehman/auth-microservice/internal/database"
	"github.com/ajaz-rehman/auth-microservice/internal/helpers"
)

type SignupRequest struct {
	FirstName string `json:"first_name" mod:"trim,lcase,title" validate:"required,alpha,ascii,min=2,max=25"`
	LastName  string `json:"last_name" mod:"trim,lcase,title" validate:"required,alpha,ascii,min=2,max=25"`
	Email     string `json:"email" mod:"trim,lcase" validate:"required,email"`
	Password  string `json:"password" validate:"required,ascii,min=8,max=50,excludes= "`
}

func SignupHandler(app *app.App) http.HandlerFunc {
	return helpers.HandleRequest(func(data SignupRequest, req *http.Request) (status int, res interface{}, err error) {
		hashedPassword, err := auth.HashPassword(data.Password)

		if err != nil {
			return
		}

		_, err = app.DB.CreateUser(req.Context(), database.CreateUserParams{
			FirstName:      data.FirstName,
			LastName:       data.LastName,
			Email:          data.Email,
			HashedPassword: hashedPassword,
		})

		if err != nil {
			return
		}

		accessToken, err := auth.CreateJWTToken(1, app.ENV.JWTSecret)

		if err != nil {
			return
		}

		refreshToken, err := auth.CreateRefreshToken()

		if err != nil {
			return
		}

		res = auth.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		return
	})
}
