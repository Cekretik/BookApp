package models

import (
	"github.com/Cekretik/BookApp/cmd/main/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name    string `gorm:"column:name" json:"name"`
	Author  string `gorm:"column:author" json:"author"`
	Release string `gorm:"column:release" json:"release"`
}

// Error implements error.
func (Book) Error() string {
	panic("unimplemented")
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	result := db.Create(&b)
	if result.Error == nil && result.RowsAffected > 0 {
		return b
	}
	return nil
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(id uint) (*Book, error) {
	var getBook Book
	db := db.Where("ID = ?", id).Find(&getBook)
	return &getBook, db.Error
}

func DeleteBook(ID uint) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(&book)
	return book
}

func UpdateBook(ID uint, updatedBook *Book) error {
	var existingBook Book
	db := db.First(&existingBook, ID)
	if db.Error != nil {
		return db.Error
	}

	if updatedBook.Name != "" {
		existingBook.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		existingBook.Author = updatedBook.Author
	}
	if updatedBook.Release != "" {
		existingBook.Release = updatedBook.Release
	}

	db = db.Save(&existingBook)
	if db.Error != nil {
		return db.Error
	}

	return nil
}
