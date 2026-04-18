package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/books", controllers.getBooks)
	router.GET("/books/:id", controllers.bookById)
	router.POST("/books", controllers.createBook)
	router.PATCH("checkout", controllers.checkoutBook)
	router.PATCH("/return", controllers.returnBook)

	router.Run("localhost:8080")
}
