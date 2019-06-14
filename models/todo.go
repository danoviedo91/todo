package models

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

//---- TODO MODEL ----//

// Todo --> struct that contains todo's fields
type Todo struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Title       string    `gorm:"type:varchar(100)"`
	Description string    `gorm:"type:varchar(255)"`
	Deadline    time.Time `gorm:"type:DATE"`
	Completed   bool      `gorm:"type:bool"`
}

//---- DATABASE CRUD OPERATIONS (METHODS) ----//

// Create inserts a row in the database
func (t Todo) Create(db *gorm.DB) {
	db.Create(&t)
}

// ReadAll returns all the tasks
func (t Todo) ReadAll(db *gorm.DB) []Todo {
	allUncompletedRecords := []Todo{}
	db.Find(&allUncompletedRecords)
	return allUncompletedRecords
}

// ReadRecord returns a single occurrence given an id
func (t Todo) ReadRecord(db *gorm.DB, id string) Todo {
	readRecord := Todo{}
	db.Where("id=?", id).First(&readRecord)
	return readRecord
}

// UpdateRecord updates a row in the database
func (t Todo) UpdateRecord(db *gorm.DB, id string) {
	db.Model(&t).Where("id = ?", id).Update("title", t.Title)
	db.Model(&t).Where("id = ?", id).Update("description", t.Description)
	db.Model(&t).Where("id = ?", id).Update("deadline", t.Deadline)
}

// UpdateCompletedRecord updates a row with boolean "Completed" true or false
func (t Todo) UpdateCompletedRecord(db *gorm.DB, id string, action string) {
	if action == "complete" {
		db.Model(&t).Where("id = ?", id).Update("completed", true)
	} else {
		db.Model(&t).Where("id = ?", id).Update("completed", false)
	}

}

// DeleteRecord removes a single record given an id
func (t Todo) DeleteRecord(db *gorm.DB, id string) {
	db.Where("id=?", id).Delete(&t)
}

//---- DATABASE UUID CREATION METHOD ----//

// BeforeCreate will set a UUID rather than numeric ID.
func (t *Todo) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}

//---- FRONT END METHODS ----//

// MonthFormatted converts type Month to type Int
func (t Todo) MonthFormatted() string {
	monthInt := int(t.Deadline.Month())
	return fmt.Sprintf("%02d", monthInt)
}

// DayFormatted converts type Month to type Int
func (t Todo) DayFormatted() string {
	return fmt.Sprintf("%02d", t.Deadline.Day())
}
