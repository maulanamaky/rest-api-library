package main

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.BookById)
	router.POST("/books", controllers.CreateBook)
	router.PATCH("/checkout", controllers.CheckoutBook)
	router.PATCH("/return", controllers.ReturnBook)

	router.Run("localhost:8080")
}
