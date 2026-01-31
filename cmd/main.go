package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"goapi/db"
	"goapi/models"
	"goapi/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file from project root
	envPath := filepath.Join(os.Getenv("PWD"), ".env")
	if err := godotenv.Load(envPath); err != nil {
		// Try loading from parent directory if running from cmd/
		if err := godotenv.Load("../.env"); err != nil {
			log.Println("Warning: Could not load .env file")
		}
	}

	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	// Auto create tables
	db.DB.AutoMigrate(
		&models.Staff{},
		&models.User{},
		&models.Todo{},
	)

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
