package db

import (
	"log"

	"github.com/danoviedo91/todo/models"
	"github.com/jinzhu/gorm"
)

// Connect establishes connection with the underlying database
func Connect() *gorm.DB {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=todo password=root sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	if !db.HasTable(&models.Todo{}) {
		db.CreateTable(&models.Todo{})
	}

	return db

}
