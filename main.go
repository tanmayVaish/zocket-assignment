package main

import (
	"log"
	"os"
	"zocket-assignment/models"
	"zocket-assignment/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Create a new gin instance
	r := gin.Default()

	// Load .env file and Create a new connection to the database
	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}
	config := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Initialize DB
	models.InitDB(config)

	// Load the routes

	// TASK 1
	routes.CrudRoute(r)

	// TASK 2
	routes.ParseCsvRoute(r)

	// Run the server
	r.Run(":8080")
}
