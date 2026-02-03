package main

import (
	"log"
	"net/http"

	"goapi/db"
	"goapi/models"
	"goapi/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Error loading .env file:", err)
	}

	// Connect to database
	if err := db.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto create tables
	db.DB.AutoMigrate(
		&models.Staff{},
		&models.User{},
		&models.Todo{},
		&models.Education{},
	)

	log.Println("Database migrated successfully")
	log.Println("Database migrated successfully")

	// Seed database
	db.Seed()
	log.Println("Database seeded successfully")

	// Register routes
	routes.RegisterRoutes()

	// Start server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
