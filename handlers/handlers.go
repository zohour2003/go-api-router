// handlers/handlers.go
package handlers

import (
	"encoding/json"
	"go-router/models"
	"net/http"

	"github.com/gorilla/mux"
)

var Messages []models.Message

func GetMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Messages)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Messages {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Message not found", http.StatusNotFound)
}
