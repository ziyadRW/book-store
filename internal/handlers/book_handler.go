package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ziyadrw/go-book-store/config"
	"github.com/ziyadrw/go-book-store/internal/models"
	"github.com/ziyadrw/go-book-store/internal/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	utils.WriteJSONResponse(w, http.StatusOK, newBooks)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid book ID"})
		return
	}

	bookDetails := models.GetBook(id)
	if bookDetails == nil {
		utils.WriteJSONResponse(w, http.StatusNotFound, map[string]string{"message": "Book not found"})
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, bookDetails)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	utils.ParseBody(r, &book)

	fmt.Printf("Received Book: %+v\n", book)

	b := book.CreateBook()
	utils.WriteJSONResponse(w, http.StatusOK, b)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid book ID"})
		return
	}
	book := models.DeleteBook(id)
	utils.WriteJSONResponse(w, http.StatusOK, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid book ID"})
		return
	}

	bookDetails := models.GetBook(id)
	if bookDetails == nil {
		utils.WriteJSONResponse(w, http.StatusNotFound, map[string]string{"message": "Book not found"})
		return
	}

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publications != "" {
		bookDetails.Publications = updateBook.Publications
	}

	db := config.GetDB()
	db.Save(&bookDetails)

	utils.WriteJSONResponse(w, http.StatusOK, bookDetails)
}
