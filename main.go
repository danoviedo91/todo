package main

import (
	"net/http"

	"github.com/danoviedo91/todo/actions"
)

func main() {

	// Serve "assets" folder as root of the server
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Handlers
	http.HandleFunc("/", actions.Index)
	http.HandleFunc("/new", actions.New)
	http.HandleFunc("/create", actions.Create)
	http.HandleFunc("/delete", actions.Delete)
	http.HandleFunc("/edit", actions.Edit)
	http.HandleFunc("/update", actions.Update)
	http.HandleFunc("/show", actions.Show)
	http.HandleFunc("/complete", actions.Complete)

	http.ListenAndServe(":8000", nil)
}
