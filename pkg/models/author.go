package models

import (
	"github.com/marcossabatini7/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name string `json:"author" gorm:"name"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Author{})
}

func (a *Author) GetAll() []Author {
	var authors []Author
	db.Find(&authors)
	return authors
}
