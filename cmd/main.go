package main

import (
	"os"

	"github.com/Thirawoot-Put/event-ticketing/order-service/internal/infrastructure/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading env file")
	}

	port := os.Getenv("PORT")
	app := server.AppServer()

	app.Start(port)
}
