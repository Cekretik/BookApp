package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Cekretik/BookApp/cmd/main/pkg/models"
	"github.com/Cekretik/BookApp/cmd/main/pkg/utils"
	"github.com/gin-gonic/gin"
)

var NewBook models.Book

func GetBook(c *gin.Context) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(res)
}

func GetBookById(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(res)
}

func CreateBook(c *gin.Context) {
	CreateBook := &models.Book{}
	utils.ParseBody(c, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(res)
}

func DeleteBook(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(res)
}

func UpdateBook(c *gin.Context) {
	var updateBook = &models.Book{}
	utils.ParseBody(c, updateBook)
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Release != "" {
		bookDetails.Release = updateBook.Release
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(res)
}
