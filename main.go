package main

import (
	"fmt"
	"os"
	"project1/bootstrap"
	"project1/database"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port: " + port)

	database.ConnetDB()
	bootstrap.Serve(":" + port)
}
