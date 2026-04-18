package repositories

import (
	"errors"
	"go-api/models"
)

var books = []models.Book{
	{ID: "1", Title: "Clean Code", Author: "Robert C. Martin", Quantity: 2},
	{ID: "2", Title: "Clean Architecture", Author: "Robert C. Martin", Quantity: 1},
	{ID: "3", Title: "The Clean Coder", Author: "Robert C. Martin", Quantity: 1},
}

func getAllBooks() []models.Book {
	return books
}

func FindBookById(id string) (*models.Book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("Book not found")
}

func createBook(newBook models.Book) models.Book {
	books = append(books, newBook)
	return newBook
}
