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

func (b *Book) GetById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func (b *Book) Create() *Book {
	db.Create(&b)
	return b
}

func (b *Book) Update(Id int64) *Book {
	db.Where("ID=?", Id).Updates(b)
	return b
}

func (b *Book) Delete(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
