package controllers

import (
	"go-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getBooks(c *gin.Context) {
	books := services.getBook()
	c.JSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := services.getBookById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.JSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books := services.addBook(newBook)
	c.JSON(http.StatusCreated, newBook)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
		return
	}

	book, err := services.checkoutBook(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.JSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
		return
	}

	book, err := services.returnBook(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.JSON(http.StatusOK, book)
}
