package server

func getEndpoints() Endpoints {
	endpoints := Endpoints{
		signupEndpoint,
	}

	return endpoints
}

var signupEndpoint = Endpoint{
	Pattern: "POST /v1/auth/signup",
	Handler: signupHandler,
	RequestBody: map[string]interface{}{
		"email":    "string",
		"password": "string",
	},
}
