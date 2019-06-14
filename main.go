package main

import (
	"log"
	"net/http"
	"os"

	"github.com/danoviedo91/todo/actions"
	"github.com/joho/godotenv"
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

	port := ""

	err := godotenv.Load()

	if err != nil {
		port = ":8000"
		log.Println(".env file not detected... using default PORT 8000")
	} else {
		port = ":" + os.Getenv("PORT")
		if port == "" {
			log.Println("PORT env variable not detected... using default PORT 8000")
			port = ":8000"
		} else {
			log.Println(".env file set with PORT", port)
		}
	}

	log.Println("Serving...")
	http.ListenAndServe(port, nil)
}
