package handlers

import (
	"net/http"
	"os"

	"github.com/ajaz-rehman/auth-microservice/internal/auth"
	"github.com/ajaz-rehman/auth-microservice/internal/helpers"
)

type SignupRequest struct {
	FirstName string `json:"first_name" mod:"trim,lcase,title" validate:"required,alpha,ascii,min=2,max=25"`
	LastName  string `json:"last_name" mod:"trim,lcase,title" validate:"required,alpha,ascii,min=2,max=25"`
	Email     string `json:"email" mod:"trim,lcase" validate:"required,email"`
	Password  string `json:"password" validate:"required,ascii,min=8,max=50,excludes= "`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	_, err := helpers.TransformAndValidateBody[SignupRequest](r.Body)

	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	accessToken, err := auth.CreateJWTToken(1, jwtSecret)

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	response := auth.Tokens{
		AccessToken:  accessToken,
		RefreshToken: "refresh",
	}

	helpers.RespondWithJSON(w, http.StatusCreated, response)
}
