package helpers

import "net/http"

type RequestHandlerFn[T any] func(data T, req *http.Request) (status int, res interface{}, err error)
