package controllers

type Controller func(data interface{}) (status int, response interface{}, err error)
