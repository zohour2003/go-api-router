package main

import (
	"go-router/handlers"
	"go-router/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	handlers.Messages = utils.FetchMessages()
	r := mux.NewRouter()
	r.HandleFunc("/api/messages", handlers.GetMessages).Methods("GET")
	r.HandleFunc("/api/messages/{id}", handlers.GetMessage).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
