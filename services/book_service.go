package services

import (
	"errors"
	"go-api/models"
	"go-api/repositories"
)

func GetBook() []models.Book {
	return repositories.GetAllBooks()
}

func GetBookById(id string) (*models.Book, error) {
	return repositories.FindBookById(id)
}

func AddBook(book models.Book) models.Book {
	return repositories.CreateBook(book)
}

func CheckoutBook(id string) (*models.Book, error) {
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

func ReturnBook(id string) (*models.Book, error) {
	book, err := repositories.FindBookById(id)

	if err != nil {
		return nil, err
	}

	book.Quantity += 1
	return book, nil
}
