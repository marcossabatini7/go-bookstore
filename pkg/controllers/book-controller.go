package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/marcossabatini7/go-bookstore/pkg/models"
	"github.com/marcossabatini7/go-bookstore/pkg/utils"
)

var Book models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := Book.GetAll()
	res, _ := json.Marshal(newBooks)

	utils.WriteResponseOk(w, res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing", err)
	}
	bookDetails, _ := Book.GetById(ID)

	res, _ := json.Marshal(bookDetails)
	utils.WriteResponseOk(w, res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.Create()

	res, _ := json.Marshal(b)
	utils.WriteResponseOk(w, res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	//	Getting vars
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	//	Create book by body data
	book := &models.Book{}
	utils.ParseBody(r, book)

	//	Updating the book by ID
	b := book.Update(ID)

	//	Writing response
	res, _ := json.Marshal(b)
	utils.WriteResponseOk(w, res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	book := Book.Delete(ID)
	res, _ := json.Marshal(book)
	utils.WriteResponseOk(w, res)
}
