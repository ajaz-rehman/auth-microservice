package tests

import (
	"os"
	"time"

	"github.com/ajaz-rehman/auth-microservice/internal/server"
)

func setupServer() {
	// Start the server
	go func() {
		err := server.ListenAndServe()

		if err != nil {
			panic("Could not start server: " + err.Error())
		}
	}()

	// Wait for the server to start
	time.Sleep(time.Second)
}

func getServerURL() string {
	return "http://localhost:" + os.Getenv("PORT")
}
