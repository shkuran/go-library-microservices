package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shkuran/go-library-microservices/book-service/book"
)

func RegisterRoutes(server *gin.Engine, book book.Handler) {
	server.GET("/books", book.GetBooks)
	server.GET("/books/:id", book.GetBookById)
	server.PUT("/books", book.UpdateAvailableCopies)
	server.POST("/books", book.AddBook)

}
