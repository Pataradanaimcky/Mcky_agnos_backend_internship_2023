package main

import (
	"log"
	"strong_password_backend/handlers"
	"strong_password_backend/services"

	"./db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db, err := db.NewDB()
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// Initialize services and handlers
	passwordService := services.NewPasswordService()
	passwordHandler := handlers.NewPasswordHandler(passwordService)

	// Set up Gin router
	router := gin.Default()

	// Strong password steps API
	router.POST("/api/strong_password_steps", passwordHandler.GetStrongPasswordSteps)

	// Start the server
	router.Run(":8080")
}
