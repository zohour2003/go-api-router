// handlers/handlers.go
package handlers

import (
	"encoding/json"
	"go-router/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Posts []models.Posts

func GetMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Fetching all messages")
	json.NewEncoder(w).Encode(Posts)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Posts {
		if item.ID == params["id"] {
			log.Printf("Fetching message with ID: %s\n", params["id"])
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	log.Printf("Message with ID: %s not found\n", params["id"])
	http.Error(w, "Message not found", http.StatusNotFound)
}
