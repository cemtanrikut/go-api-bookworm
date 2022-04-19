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

func Add(book Book) error {
	return nil
}

func Get(bookName string) Book {
	return Book{}
}

func Delete(bookID string) error {
	return nil
}
