package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// template.ParseFiles parses the index.html file in root
// template.Must so code panics
var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("<h1>Hello World!</h1>"))
	tpl.Execute(w, nil)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// attempt to set the port variable based on the PORT environment variable
	port := os.Getenv("PORT")
	// Getenv returns empty string if environment variable not present
	if port == "" {
		port = "3000"
	}

	// creates an HTTP request multiplexer "mux"
	mux := http.NewServeMux()

	// HandleFunc variables:
	// 1 - pattern string
	// 2 - func(w http.ResponseWriter, r *http.Request)
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}