package book

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shkuran/go-library-microservices/book-service/utils"
)

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) Handler {
	return Handler{repo: repo}
}

func (h Handler) GetBooks(context *gin.Context) {
	books, err := h.repo.GetAll()
	if err != nil {
		utils.HandleInternalServerError(context, "Could not fetch books!", err)
		return
	}
	context.JSON(http.StatusOK, books)
}

func (h Handler) GetBookById(context *gin.Context) {
	bookId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		utils.HandleBadRequest(context, "Could not parse bookId!", err)
		return
	}

	book, err := h.repo.GetById(bookId)
	if err != nil {
		utils.HandleInternalServerError(context, "Could not fetch the book!", err)
		return
	}
	context.JSON(http.StatusOK, book)
}

func (h Handler) UpdateAvailableCopies(context *gin.Context) {
	var updateCopies struct {
		BookID          int64 `json:"book_id"`
		AvailableCopies int64 `json:"available_copies"`
	}

	err := context.ShouldBindJSON(&updateCopies)
	if err != nil {
		utils.HandleBadRequest(context, "Could not parse request data!", err)
		return
	}

	err = h.repo.UpdateAvailableCopies(updateCopies.BookID, updateCopies.AvailableCopies)
	if err != nil {
		utils.HandleInternalServerError(context, "Could not update the number of book copies!", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Number of copies updated successfully"})
}

func (h Handler) AddBook(context *gin.Context) {
	var b Book
	err := context.ShouldBindJSON(&b)
	if err != nil {
		utils.HandleBadRequest(context, "Could not parse request data!", err)
		return
	}
	err = h.repo.Save(b)
	if err != nil {
		utils.HandleInternalServerError(context, "Could not add book!", err)
		return
	}
	utils.HandleStatusCreated(context, "Book added!")
}
