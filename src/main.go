package main

import (
	"log"
	"os"

	"example.com/go-poc/database"
	"example.com/go-poc/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	port := os.Getenv("SERVER_PORT")
	connectionString := os.Getenv("CONNECTION_STRING")

	database.DatabaseConnection(connectionString)
	routes.HandleRequest(port)
}