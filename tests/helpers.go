package tests

import (
	"os"
	"time"

	"github.com/ajaz-rehman/auth-microservice/internal/server"
	"github.com/joho/godotenv"
)

func setupServer() {
	// Start the server
	go func() {
		godotenv.Load()

		port := os.Getenv("PORT")

		err := server.ListenAndServe(port)

		if err != nil {
			panic("Could not start server: " + err.Error())
		}
	}()

	// Wait for the server to start
	time.Sleep(time.Second)
}
