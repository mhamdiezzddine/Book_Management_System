package models

import (
	"github.com/jinzhu/gorm"

	"github.com/mhamdiezzddine/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

//intialize the database

func init() {

	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})

}

// CreateBook function
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// Get ALl Books function
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// Get Book by id function
func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
}

// Delete book function
func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return book
}
