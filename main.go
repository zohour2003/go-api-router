package main

import (
	"go-router/handlers"
	"go-router/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting application")
	handlers.Posts = utils.FetchMessages()

	r := mux.NewRouter()
	r.HandleFunc("/api/posts", handlers.GetMessages).Methods("GET")
	r.HandleFunc("/api/posts/{userId}", handlers.GetMessage).Methods("GET")

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
