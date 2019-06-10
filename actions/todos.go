package actions

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
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
	r.ParseForm()
	todo := models.Todo{}

	layout := "2006-01-02"
	duedate, err := time.Parse(layout, r.FormValue("todo-date"))

	if err != nil {
		fmt.Println(err)
	}

	todo.Deadline = duedate
	todo.Description = r.FormValue("todo-description")
	todo.Title = r.FormValue("todo-title")
	todo.Completed = false

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

	todo := models.Todo{}

	layout := "01/02/2006"
	duedateDecoded, _ := url.QueryUnescape(r.URL.Query()["limitdate"][0])
	duedate, err := time.Parse(layout, duedateDecoded)

	if err != nil {
		fmt.Println(err)
	}

	todo.Title, _ = url.QueryUnescape(r.URL.Query()["title"][0])
	todo.Description, _ = url.QueryUnescape(r.URL.Query()["description"][0])
	todo.Deadline = duedate

	//Parse HTML template
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

	todo := models.Todo{}

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
