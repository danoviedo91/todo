package db

import (
	"log"

	"github.com/danoviedo91/todo/models"
	"github.com/jinzhu/gorm"
)

// Connect establishes connection with the underlying database
func Connect() *gorm.DB {

	db, err := gorm.Open("postgres", "dbname=d8gqvbans6jeot host=ec2-174-129-18-42.compute-1.amazonaws.com port=5432 user=toxacrziddmgkh password=a1ee883bfdbbb2dd788c8f9b19a9da4df0a4166dd3357907bd9d821435ddde91 sslmode=require")

	if err != nil {
		log.Fatal(err)
	}

	if !db.HasTable(&models.Todo{}) {
		db.CreateTable(&models.Todo{})
	}

	return db

}
