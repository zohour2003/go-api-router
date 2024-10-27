// utils/fetch.go
package utils

import (
	"encoding/json"
	"go-router/models"
	"io/ioutil"
	"log"
	"net/http"
)

func FetchMessages() []models.Posts {
	log.Println("Fetching messages from JSONPlaceholder . . .")
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var messages []models.Posts
	json.Unmarshal(body, &messages)
	log.Printf("Fetched %d messages\n", len(messages))
	return messages
}
