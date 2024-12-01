package server

import (
	"net/http"
	"os"
)

type SignupRequestBody struct {
	FirstName string `json:"first_name" mod:"trim,lcase,title" validate:"required,alpha,ascii,min=2,max=25"`
	LastName  string `json:"last_name" mod:"trim,lcase,title" validate:"required,alpha,ascii,min=2,max=25"`
	Email     string `json:"email" mod:"trim,lcase" validate:"required,email"`
	Password  string `json:"password" validate:"required,ascii,min=8,max=50,excludes= "`
}

func signup(data SignupRequestBody) (status int, response any, err error) {
	status = http.StatusCreated

	jwtSecret := os.Getenv("JWT_SECRET")

	accessToken, err := MakeJWT(1, jwtSecret)

	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	response = Tokens{
		AccessToken:  accessToken,
		RefreshToken: "refresh",
	}
	return
}
