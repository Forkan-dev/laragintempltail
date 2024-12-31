package main

import (
	"fmt"
	"os"
	"project1/bootstrap"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// GET PORT FROM ENV
	port := os.Getenv("PORT")
	fmt.Println("PORT: " + port)
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port: " + port)

	bootstrap.Serve(":" + port)
}
