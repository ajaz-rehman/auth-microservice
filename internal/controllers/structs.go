package controllers

import "net/http"

type Controller func(resp *http.Request) (status int, response interface{}, err error)
