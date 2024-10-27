package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"go-router/models"
)

func FetchMessages() []models.Message {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var messages []models.Message
	json.Unmarshal(body, &messages)
	return messages
}
