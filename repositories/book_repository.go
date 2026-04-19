package repositories

import (
	"errors"
	"go-api/models"
)

var books = []models.Book{
	{ID: "1", Title: "Clean Code", Author: "Robert C. Martin", Quantity: 2},
	{ID: "2", Title: "Clean Architecture", Author: "Robert C. Martin", Quantity: 5},
	{ID: "3", Title: "The Maintenance Coder", Author: "Robert C. Martin", Quantity: 3},
}

func GetAllBooks() []models.Book {
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

func CreateBook(newBook models.Book) models.Book {
	books = append(books, newBook)
	return newBook
}
