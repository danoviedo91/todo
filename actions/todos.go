package actions

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	db "github.com/danoviedo91/todo/database"
	"github.com/danoviedo91/todo/models"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

// Index is used to parse index.html the first time user enters the website
func Index(w http.ResponseWriter, r *http.Request) {

	// queryValues gets all GET parameters sent with the URL
	queryValues := r.URL.Query()

	// Establish database connection
	db := db.Connect()
	defer db.Close()

	// Declare empty struct for giving db the needed structure
	todo := models.Todo{}

	// Initialize database-query variables

	allRecords := todo.ReadAll(db)
	records := []models.Todo{}
	incompletedTasks := 0
	filterStatus := struct {
		Incompleted bool
		Completed   bool
	}{
		false,
		false,
	}

	// If /?completed=true

	if queryValues.Get("status") == "completed" {
		for _, row := range allRecords {
			if row.Completed == true {
				records = append(records, row)
			}
		}
		incompletedTasks = len(allRecords) - len(records)
		filterStatus.Completed = true
		// If /?completed=false
	} else if queryValues.Get("status") == "incompleted" {
		for _, row := range allRecords {
			if row.Completed == false {
				records = append(records, row)
			}
		}
		incompletedTasks = len(records)
		filterStatus.Incompleted = true
		// If /
	} else {
		for _, row := range allRecords {
			if row.Completed == false {
				records = append(records, row)
			}
		}
		incompletedTasks = len(records)
		records = allRecords
	}

	// Prepare to send data to template

	templateData := struct {
		PendingTasksNumber int
		CurrentDateTime    time.Time
		TasksArray         []models.Todo
		TaskStruct         models.Todo
		FilterStatus       struct {
			Incompleted bool
			Completed   bool
		}
	}{
		PendingTasksNumber: incompletedTasks,
		CurrentDateTime:    time.Now(),
		TasksArray:         records,
		FilterStatus:       filterStatus,
	}

	// Parse HTML template
	html, err := template.ParseFiles("templates/todos/index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = html.Execute(w, templateData)
	if err != nil {
		log.Fatal(err)
	}

}

// New parses new.html which contains form submission
func New(w http.ResponseWriter, r *http.Request) {

	// queryValues gets all GET parameters sent with the URL
	queryValues := r.URL.Query()

	// Establish database connection
	db := db.Connect()
	defer db.Close()

	// Declare empty struct for giving db the needed structure
	todo := models.Todo{}

	// Initialize database-query variables

	allRecords := todo.ReadAll(db)
	records := []models.Todo{}
	incompletedTasks := 0

	for _, row := range allRecords {
		if row.Completed == false {
			records = append(records, row)
		}
	}
	incompletedTasks = len(records)

	filterStatus := struct {
		Incompleted bool
		Completed   bool
	}{
		false,
		false,
	}

	if queryValues.Get("status") == "completed" {
		// If /?completed=true
		filterStatus.Completed = true
	} else if queryValues.Get("status") == "incompleted" {
		// If /?completed=false
		filterStatus.Incompleted = true
	}

	// Prepare to send data to template

	templateData := struct {
		PendingTasksNumber int
		CurrentDateTime    time.Time
		TasksArray         []models.Todo
		TaskStruct         models.Todo
		FilterStatus       struct {
			Incompleted bool
			Completed   bool
		}
	}{
		PendingTasksNumber: incompletedTasks,
		CurrentDateTime:    time.Now(),
		FilterStatus:       filterStatus,
	}

	// Parse HTML template
	html, err := template.ParseFiles("templates/todos/new.html")
	if err != nil {
		log.Fatal(err)
	}

	err = html.Execute(w, templateData)
	if err != nil {
		log.Fatal(err)
	}

}

// Create fills the struct with info and returns it to index.html
func Create(w http.ResponseWriter, r *http.Request) {

	// Assign to r the POST information sent with the form
	r.ParseForm()

	// Establish database connection
	db := db.Connect()
	defer db.Close()

	// Get Time.time correct format for todo.Deadline
	layout := "2006-01-02"
	duedate, err := time.Parse(layout, r.FormValue("todo-date"))

	if err != nil {
		fmt.Println(err)
	}

	// Declare empty struct for giving db the needed structure
	todo := models.Todo{}

	// Assign values to todo struct
	todo.Deadline = duedate
	todo.Description = r.FormValue("todo-description")
	todo.Title = r.FormValue("todo-title")
	todo.Completed, err = strconv.ParseBool(r.FormValue("todo-completed"))

	// Insert record into the database
	todo.Create(db)

	// Initialize database-query variables

	allRecords := todo.ReadAll(db)
	records := []models.Todo{}
	incompletedTasks := 0

	for _, row := range allRecords {
		if row.Completed == false {
			records = append(records, row)
		}
	}
	incompletedTasks = len(records)
	taskToBeShowed := todo

	// Prepare to send data to template

	templateData := struct {
		PendingTasksNumber int
		CurrentDateTime    time.Time
		TasksArray         []models.Todo
		TaskStruct         models.Todo
		FilterStatus       struct {
			Incompleted bool
			Completed   bool
		}
	}{
		PendingTasksNumber: incompletedTasks,
		CurrentDateTime:    time.Now(),
		TaskStruct:         taskToBeShowed,
	}

	//Parse HTML template
	html, err := template.ParseFiles("templates/todos/show.html")
	if err != nil {
		log.Fatal(err)
	}

	err = html.Execute(w, templateData)
	if err != nil {
		log.Fatal(err)
	}

}

// Delete removes the record
func Delete(w http.ResponseWriter, r *http.Request) {

	// queryValues gets all GET parameters sent with the URL
	queryValues := r.URL.Query()

	// Catch the id
	id := queryValues.Get("id")

	// Establish database connection
	db := db.Connect()
	defer db.Close()

	// Declare empty struct for giving db the needed structure
	todo := models.Todo{}

	// Delete record given the id
	todo.DeleteRecord(db, id)

	// Initialize database-query variables

	allRecords := todo.ReadAll(db)
	records := []models.Todo{}
	incompletedTasks := 0
	filterStatus := struct {
		Incompleted bool
		Completed   bool
	}{
		false,
		false,
	}

	// If /?completed=true

	if queryValues.Get("status") == "completed" {
		for _, row := range allRecords {
			if row.Completed == true {
				records = append(records, row)
			}
		}
		incompletedTasks = len(allRecords) - len(records)
		filterStatus.Completed = true
		// If /?completed=false
	} else if queryValues.Get("status") == "incompleted" {
		for _, row := range allRecords {
			if row.Completed == false {
				records = append(records, row)
			}
		}
		incompletedTasks = len(records)
		filterStatus.Incompleted = true
		// If /
	} else {
		for _, row := range allRecords {
			if row.Completed == false {
				records = append(records, row)
			}
		}
		incompletedTasks = len(records)
		records = allRecords
	}

	// Prepare to send data to template

	templateData := struct {
		PendingTasksNumber int
		CurrentDateTime    time.Time
		TasksArray         []models.Todo
		TaskStruct         models.Todo
		FilterStatus       struct {
			Incompleted bool
			Completed   bool
		}
	}{
		PendingTasksNumber: incompletedTasks,
		CurrentDateTime:    time.Now(),
		TasksArray:         records,
		FilterStatus:       filterStatus,
	}

	// Parse HTML template
	html, err := template.ParseFiles("templates/todos/index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = html.Execute(w, templateData)
	if err != nil {
		log.Fatal(err)
	}
}

// Complete marks item as completed or incomplete
func Complete(w http.ResponseWriter, r *http.Request) {

	// queryValues gets all GET parameters sent with the URL
	queryValues := r.URL.Query()

	// Catch the id
	id := r.FormValue("id")

	// Declare empty struct for giving db the needed structure
	todo := models.Todo{}

	// Establish database connection
	db := db.Connect()
	defer db.Close()

	// Catch the struct to be edited
	todo = todo.ReadRecord(db, id)

	// Update record into the database
	todo.UpdateCompletedRecord(db, id)

	// Initialize database-query variables

	allRecords := todo.ReadAll(db)
	records := []models.Todo{}
	incompletedTasks := 0
	filterStatus := struct {
		Incompleted bool
		Completed   bool
	}{
		false,
		false,
	}

	// If /?completed=true

	if queryValues.Get("status") == "completed" {
		for _, row := range allRecords {
			if row.Completed == true {
				records = append(records, row)
			}
		}
		incompletedTasks = len(allRecords) - len(records)
		filterStatus.Completed = true
		// If /?completed=false
	} else if queryValues.Get("status") == "incompleted" {
		for _, row := range allRecords {
			if row.Completed == false {
				records = append(records, row)
			}
		}
		incompletedTasks = len(records)
		filterStatus.Incompleted = true
		// If /
	} else {
		for _, row := range allRecords {
			if row.Completed == false {
				records = append(records, row)
			}
		}
		incompletedTasks = len(records)
		records = allRecords
	}

	// Prepare to send data to template

	templateData := struct {
		PendingTasksNumber int
		CurrentDateTime    time.Time
		TasksArray         []models.Todo
		TaskStruct         models.Todo
		FilterStatus       struct {
			Incompleted bool
			Completed   bool
		}
	}{
		PendingTasksNumber: incompletedTasks,
		CurrentDateTime:    time.Now(),
		TasksArray:         records,
		FilterStatus:       filterStatus,
	}

	// Parse HTML template
	html, err := template.ParseFiles("templates/todos/index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = html.Execute(w, templateData)
	if err != nil {
		log.Fatal(err)
	}
}

// Edit allows changing the task information
func Edit(w http.ResponseWriter, r *http.Request) {

	// queryValues gets all GET parameters sent with the URL
	queryValues := r.URL.Query()

	// Establish database connection
	db := db.Connect()
	defer db.Close()

	// Declare empty struct for giving db the needed structure
	todo := models.Todo{}

	// Catch the id
	id := queryValues.Get("id")

	// Catch the struct to be edited
	taskToBeEdited := todo.ReadRecord(db, id)

	// Initialize database-query variables

	allRecords := todo.ReadAll(db)
	records := []models.Todo{}
	incompletedTasks := 0

	for _, row := range allRecords {
		if row.Completed == false {
			records = append(records, row)
		}
	}
	incompletedTasks = len(records)

	filterStatus := struct {
		Incompleted bool
		Completed   bool
	}{
		false,
		false,
	}

	if queryValues.Get("status") == "completed" {
		// If /?completed=true
		filterStatus.Completed = true
	} else if queryValues.Get("status") == "incompleted" {
		// If /?completed=false
		filterStatus.Incompleted = true
	}

	// Prepare to send data to template

	templateData := struct {
		PendingTasksNumber int
		CurrentDateTime    time.Time
		TasksArray         []models.Todo
		TaskStruct         models.Todo
		FilterStatus       struct {
			Incompleted bool
			Completed   bool
		}
	}{
		PendingTasksNumber: incompletedTasks,
		CurrentDateTime:    time.Now(),
		TaskStruct:         taskToBeEdited,
		FilterStatus:       filterStatus,
	}

	// Parse HTML template
	html, err := template.ParseFiles("templates/todos/edit.html")
	if err != nil {
		log.Fatal(err)
	}

	err = html.Execute(w, templateData)
	if err != nil {
		log.Fatal(err)
	}

}

// Update fills the struct with info and returns it to index.html
func Update(w http.ResponseWriter, r *http.Request) {

	// Assign to r the POST information sent with the form
	r.ParseForm()

	// Establish database connection
	db := db.Connect()
	defer db.Close()

	// Declare empty struct for giving db the needed structure
	todo := models.Todo{}

	// Get Time.time correct format for todo.Deadline
	layout := "2006-01-02"
	duedate, err := time.Parse(layout, r.FormValue("todo-date"))

	if err != nil {
		fmt.Println(err)
	}

	// Assign values to todo struct
	id := r.FormValue("todo-id")
	todo.Deadline = duedate
	todo.Description = r.FormValue("todo-description")
	todo.Title = r.FormValue("todo-title")
	todo.Completed, err = strconv.ParseBool(r.FormValue("todo-completed"))

	// Update record into the database
	todo.UpdateRecord(db, id)

	// Initialize database-query variables

	allRecords := todo.ReadAll(db)
	records := []models.Todo{}
	incompletedTasks := 0

	for _, row := range allRecords {
		if row.Completed == false {
			records = append(records, row)
		}
	}
	incompletedTasks = len(records)
	taskToBeShowed := todo

	// Prepare to send data to template

	templateData := struct {
		PendingTasksNumber int
		CurrentDateTime    time.Time
		TasksArray         []models.Todo
		TaskStruct         models.Todo
		FilterStatus       struct {
			Incompleted bool
			Completed   bool
		}
	}{
		PendingTasksNumber: incompletedTasks,
		CurrentDateTime:    time.Now(),
		TaskStruct:         taskToBeShowed,
	}

	//Parse HTML template
	html, err := template.ParseFiles("templates/todos/show.html")
	if err != nil {
		log.Fatal(err)
	}

	err = html.Execute(w, templateData)
	if err != nil {
		log.Fatal(err)
	}

}

// Show marks item as completed or incomplete
func Show(w http.ResponseWriter, r *http.Request) {

	// queryValues gets all GET parameters sent with the URL
	queryValues := r.URL.Query()

	// Establish database connection
	db := db.Connect()
	defer db.Close()

	// Catch the id
	id := queryValues.Get("id")

	// Declare empty struct for giving db the needed structure
	todo := models.Todo{}

	// Catch the struct to be edited
	taskToBeShowed := todo.ReadRecord(db, id)

	// Initialize database-query variables

	allRecords := todo.ReadAll(db)
	records := []models.Todo{}
	incompletedTasks := 0

	for _, row := range allRecords {
		if row.Completed == false {
			records = append(records, row)
		}
	}
	incompletedTasks = len(records)

	filterStatus := struct {
		Incompleted bool
		Completed   bool
	}{
		false,
		false,
	}

	if queryValues.Get("status") == "completed" {
		// If /?completed=true
		filterStatus.Completed = true
	} else if queryValues.Get("status") == "incompleted" {
		// If /?completed=false
		filterStatus.Incompleted = true
	}

	// Prepare to send data to template

	templateData := struct {
		PendingTasksNumber int
		CurrentDateTime    time.Time
		TasksArray         []models.Todo
		TaskStruct         models.Todo
		FilterStatus       struct {
			Incompleted bool
			Completed   bool
		}
	}{
		PendingTasksNumber: incompletedTasks,
		CurrentDateTime:    time.Now(),
		TaskStruct:         taskToBeShowed,
		FilterStatus:       filterStatus,
	}

	// Parse HTML template
	html, err := template.ParseFiles("templates/todos/show.html")
	if err != nil {
		log.Fatal(err)
	}

	err = html.Execute(w, templateData)
	if err != nil {
		log.Fatal(err)
	}
}
