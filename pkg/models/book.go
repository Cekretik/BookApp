package models

import (
	"github.com/Cekretik/BookApp/cmd/main/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.model
	Name    string `json:"name"`
	Author  string `json:"author"`
	Release int    `json:"release"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
