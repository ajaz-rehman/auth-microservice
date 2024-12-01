package controllers

import (
	"net/http"
	"os"

	"github.com/ajaz-rehman/auth-microservice/internal/auth"
)

type SignupRequest struct {
	FirstName string `json:"first_name" mod:"trim,lcase,title" validate:"required,alpha,ascii,min=2,max=25"`
	LastName  string `json:"last_name" mod:"trim,lcase,title" validate:"required,alpha,ascii,min=2,max=25"`
	Email     string `json:"email" mod:"trim,lcase" validate:"required,email"`
	Password  string `json:"password" validate:"required,ascii,min=8,max=50,excludes= "`
}

func Signup(data SignupRequest) (status int, response interface{}, err error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	accessToken, err := auth.CreateJWTToken(1, jwtSecret)

	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	status = http.StatusCreated
	response = auth.Tokens{
		AccessToken:  accessToken,
		RefreshToken: "refresh",
	}

	return
}
