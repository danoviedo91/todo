package db

import (
	"log"
	"os"

	"github.com/danoviedo91/todo/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// Connect establishes connection with the underlying database
func Connect() *gorm.DB {

	_ = godotenv.Load()

	databaseURL := os.Getenv("DATABASE_URL")

	db, err := gorm.Open("postgres", databaseURL)

	if err != nil {
		log.Fatal(err)
	}

	if !db.HasTable(&models.Todo{}) {
		db.CreateTable(&models.Todo{})
	}

	return db

}
