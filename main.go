package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ln0rd/api-golang-persons/models"
	"github.com/ln0rd/api-golang-persons/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	httpPort := os.Getenv("HTTP_PORT")

	models.Personalities = []models.Personality{
		{ID: 1, Name: "John Doe", History: "A fictional character often used as a placeholder."},
		{ID: 2, Name: "Jane Smith", History: "A common name used in various contexts."},
		{ID: 3, Name: "Alice Johnson", History: "A name often associated with curiosity and adventure."},
	}

	fmt.Println("Initializing the application...")
	routes.HandleRequest(httpPort)
}
