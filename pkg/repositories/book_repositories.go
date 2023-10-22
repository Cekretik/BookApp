package repositories

import (
	"net/http"
	"strconv"

	"github.com/Cekretik/BookApp/cmd/main/pkg/models"
	"github.com/Cekretik/BookApp/cmd/main/pkg/utils"
	"github.com/gin-gonic/gin"
)

var NewBook models.Book

func GetBook(c *gin.Context) {
	newBooks := models.GetAllBooks()
	if newBooks == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get books"})
		return
	}
	c.JSON(http.StatusOK, newBooks)
}

func GetBookById(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID", "details": err.Error()})
		return
	}

	bookDetails, err := models.GetBookById(uint(ID))

	c.JSON(http.StatusOK, bookDetails)
}

func CreateBook(c *gin.Context) {
	createBook := &models.Book{}
	if err := c.ShouldBindJSON(createBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided", "details": err.Error()})
		return
	}
	createdBook := createBook.CreateBook()
	if createdBook == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}
	c.JSON(http.StatusOK, createdBook)
}

func DeleteBook(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid ID", err)
		return
	}

	err = models.DeleteBook(uint(ID))
	respondWithMessage(c, http.StatusOK, "Book deleted successfully")
}

func UpdateBook(c *gin.Context) {
	var requestBook models.Book
	utils.ParseBody(c, &requestBook)

	bookId := c.Param("bookId")
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid ID", err)
		return
	}
	_, err = models.GetBookById(uint(ID))
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to update book", err)
		return
	}

	respondWithMessage(c, http.StatusOK, "Book updated successfully")
}
