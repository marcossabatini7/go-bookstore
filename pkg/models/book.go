package models

import (
	"github.com/marcossabatini7/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name" gorm:"name"`
	Author      string `json:"author" gorm:"author"`
	Publication string `json:"publication" gorm:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) GetAll() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func (b *Book) UpdateBook(Id int64) *Book {
	db.Where("ID=?", Id).Updates(b)
	return b
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
