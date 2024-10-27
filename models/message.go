package models

type Posts struct {
	UserId int    `json:"userId"`
	ID     string `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
