package main

import (
	"log"
	"net/http"
	"github.com/andymartinezot/crud-app-go/backend/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Initiate)
	http.HandleFunc("/create", handlers.Create)
	http.HandleFunc("/insert", handlers.Insert)
	http.HandleFunc("/update", handlers.Update)
	http.HandleFunc("/edit", handlers.Edit)
	http.HandleFunc("/delete", handlers.Delete)

	log.Println("server up and running...")

	// Initiate server
	http.ListenAndServe(":8000", nil)
}