package controllers

type Controller[T interface{}] func(data T) (status int, response interface{}, err error)
