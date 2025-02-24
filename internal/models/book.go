package models

import (
	"github.com/ziyadrw/go-book-store/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name         string `json:"name"`
	Author       string `json:"author"`
	Publications string `json:"publications"`
}

func init() {
	config.Connect()
	db = config.GetDB()

	if err := db.AutoMigrate(&Book{}); err != nil {
		panic(err)
	}
}

func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBook(Id int64) *Book {
	var book Book
	result := db.Where("ID = ?", Id).First(&book)
	if result.Error != nil {
		return nil
	}
	return &book
}

func DeleteBook(Id int64) *Book {
	var book Book
	if err := db.Where("ID = ?", Id).First(&book).Error; err != nil {
		return nil
	}
	db.Delete(&book)
	return &book
}
