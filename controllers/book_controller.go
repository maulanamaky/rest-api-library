package controllers

import (
	"go-api/models"
	"go-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	books := services.GetBook()
	c.JSON(http.StatusOK, books)
}

func BookById(c *gin.Context) {
	id := c.Param("id")
	book, err := services.GetBookById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books := services.AddBook(newBook)
	c.JSON(http.StatusCreated, books)
}

func CheckoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
		return
	}

	book, err := services.CheckoutBook(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.JSON(http.StatusOK, book)
}

func ReturnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
		return
	}

	book, err := services.ReturnBook(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.JSON(http.StatusOK, book)
}
