package actions

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/danoviedo91/todo/models"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

// Index is used to parse index.html the first time user enters the website
func Index(w http.ResponseWriter, r *http.Request) {

	// Parse HTML template
	html, err := template.ParseFiles("templates/todos/index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = html.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}

}

// New parses new.html which contains form submission
func New(w http.ResponseWriter, r *http.Request) {

	// Parse HTML template
	html, err := template.ParseFiles("templates/todos/new.html")
	if err != nil {
		log.Fatal(err)
	}

	err = html.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}

}

// Create fills the struct with info and returns it to index.html
func Create(w http.ResponseWriter, r *http.Request) {

	// Assign to r the POST information sent with the form
	r.ParseForm()

	// Declare empty todo struct
	todo := models.Todo{}

	// Get Time.time correct format for todo.Deadline
	layout := "2006-01-02"
	duedate, err := time.Parse(layout, r.FormValue("todo-date"))

	if err != nil {
		fmt.Println(err)
	}

	// Assign values to todo struct
	todo.Deadline = duedate
	todo.Description = r.FormValue("todo-description")
	todo.Title = r.FormValue("todo-title")
	todo.Completed, err = strconv.ParseBool(r.FormValue("todo-completed"))

	if err != nil {
		fmt.Println(err)
	}

	//Establish database connection

	// db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=todo password=root")
	// defer db.Close()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	//Parse HTML template
	html, err := template.ParseFiles("templates/todos/index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = html.Execute(w, todo)
	if err != nil {
		log.Fatal(err)
	}
}

// Edit allows changing the task information
func Edit(w http.ResponseWriter, r *http.Request) {

	// queryValues gets all GET parameters sent with the URL
	queryValues := r.URL.Query()

	// Assign true or false to todo.Completed
	completed, _ := strconv.ParseBool(queryValues.Get("completed"))

	todo := models.Todo{
		Completed: completed,
	}

	// Get Time.time correct format for todo.Deadline
	layout := "01/02/2006"
	duedateDecoded, _ := url.QueryUnescape(queryValues.Get("limitdate"))
	duedate, err := time.Parse(layout, duedateDecoded)

	if err != nil {
		fmt.Println(err)
	}

	// Assign values to todo struct

	todo.Title, _ = url.QueryUnescape(queryValues.Get("title"))
	todo.Description, _ = url.QueryUnescape(queryValues.Get("description"))
	todo.Deadline = duedate

	// Parse HTML template
	html, err := template.ParseFiles("templates/todos/edit.html")
	if err != nil {
		log.Fatal(err)
	}

	err = html.Execute(w, todo)
	if err != nil {
		log.Fatal(err)
	}

}

// Delete removes the record
func Delete(w http.ResponseWriter, r *http.Request) {

	// todo gets cleared out
	todo := models.Todo{}

	// Parse HTML template
	html, err := template.ParseFiles("templates/todos/index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = html.Execute(w, todo)
	if err != nil {
		log.Fatal(err)
	}
}

// Completed marks item as completed or incomplete
func Completed(w http.ResponseWriter, r *http.Request) {

	// queryValues gets all GET parameters sent with the URL
	queryValues := r.URL.Query()

	// Assign true or false to todo.Completed
	completed, _ := strconv.ParseBool(queryValues.Get("completed"))

	todo := models.Todo{
		Completed: completed,
	}

	// Get Time.time correct format for todo.Deadline

	layout := "01/02/2006"
	duedateDecoded, _ := url.QueryUnescape(queryValues.Get("limitdate"))
	duedate, err := time.Parse(layout, duedateDecoded)

	if err != nil {
		fmt.Println(err)
	}

	// Assign values to todo struct
	todo.Title, _ = url.QueryUnescape(queryValues.Get("title"))
	todo.Description, _ = url.QueryUnescape(queryValues.Get("description"))
	todo.Deadline = duedate

	// Parse HTML template
	html, err := template.ParseFiles("templates/todos/index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = html.Execute(w, todo)
	if err != nil {
		log.Fatal(err)
	}
}
