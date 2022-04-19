package book

import (
	userdata "github.com/cemtanrikut/go-api-bookworm/api/user"
)

type Book struct {
	BookID      string        `json:"book_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Category    string        `json:"category"`
	ReleaseDate string        `json:"release_date"`
	Author      userdata.User `json:"author"`
}
