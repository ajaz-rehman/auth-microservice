package handlers

import (
	"net/http"
	"strings"

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
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := helpers.TransformAndValidateBody[SignupRequest](r.Body)

		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, err)
		}

		hashedPassword, err := auth.HashPassword(data.Password)

		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		_, err = app.DB.CreateUser(r.Context(), database.CreateUserParams{
			FirstName:      data.FirstName,
			LastName:       data.LastName,
			Email:          data.Email,
			HashedPassword: hashedPassword,
		})

		if err != nil {
			if strings.Contains(err.Error(), "unique constraint") {
				helpers.RespondWithError(w, http.StatusConflict, err)
				return
			}

			helpers.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		accessToken, err := auth.CreateJWTToken(1, app.ENV.JWTSecret)

		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		refreshToken, err := auth.CreateRefreshToken()

		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		response := auth.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		helpers.RespondWithJSON(w, http.StatusCreated, response)
	}
}
