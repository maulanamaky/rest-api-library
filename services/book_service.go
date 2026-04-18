package services

import (
	"errors"
	"go-api/models"
	"go-api/repositories"
)

func getBook() []models.Book {
	return repositories.getAllBooks()
}

func getBookById(id string) (*models.Book, error) {
	return repositories.FindBookById(id)
}

func addBook(book models.Book) models.Book {
	return repositories.createBook(book)
}

func checkoutBook(id string) (*models.Book, error) {
	book, err := repositories.FindBookById(id)

	if err != nil {
		return nil, err
	}

	if book.Quantity <= 0 {
		return nil, errors.New("book not available")
	}

	book.Quantity -= 1
	return book, nil
}

func returnBook(id string) (*models.Book, error) {
	book, err := repositories.FindBookById(id)

	if err != nil {
		return nil, err
	}

	book.Quantity += 1
	return book, nil
}
