package main

import (
	"log"
	"net/http"
	"os"

	"github.com/danoviedo91/todo/actions"
	"github.com/joho/godotenv"
)

func main() {

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", actions.Index)
	http.HandleFunc("/new", actions.New)
	http.HandleFunc("/create", actions.Create)
	http.HandleFunc("/delete", actions.Delete)
	http.HandleFunc("/edit", actions.Edit)
	http.HandleFunc("/update", actions.Update)
	http.HandleFunc("/show", actions.Show)
	http.HandleFunc("/complete", actions.Complete)

	port := ""

	_ = godotenv.Load()

	port = ":" + os.Getenv("PORT")

	log.Println("Serving...")
	http.ListenAndServe(port, nil)
}
