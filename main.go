package main

import (
	"Mcky_agnos_backend_internship_2023/api"
	"Mcky_agnos_backend_internship_2023/db"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("DATABASE_URL:", os.Getenv("DATABASE_URL"))

	db.InitDatabase()
	connStr := os.Getenv("DATABASE_URL")
	fmt.Println("Connection string:", connStr)

	fmt.Println("Creating logs table...")
	err = db.CreateLogsTable()
	if err != nil {
		log.Fatal("Failed to create logs table: ", err)
	}
	fmt.Println("Logs table created successfully.")

	router := gin.Default()

	router.Use(api.RequestLogger())

	api.SetupRoutes(router)

	errServer := router.Run(":8080")
	if errServer != nil {
		log.Fatal("Failed to start the server: ", errServer)
	}
}
