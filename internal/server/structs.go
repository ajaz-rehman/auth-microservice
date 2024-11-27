package server

import "net/http"

type Endpoints map[string]http.HandlerFunc
