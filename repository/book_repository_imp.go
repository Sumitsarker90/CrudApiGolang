package repository

import (
	request "bookapi/data/request/response"
	"bookapi/helper"
	"bookapi/model"
	"errors"

	"gorm.io/gorm"
)

type BookrespositoryImpl struct {
	Db *gorm.DB
}

func NewBookrepo(Db *gorm.DB) Bookrespository {
	return &BookrespositoryImpl{Db: Db}
}

func (b *BookrespositoryImpl) Delete(booksId int) {
	var books model.Book

	result := b.Db.Where("id= ?", booksId).Delete(&books)

	helper.PanicError(result.Error)

}

func (b *BookrespositoryImpl) Findall() []model.Book {
	var books []model.Book
	result := b.Db.Find(&books)
	helper.PanicError(result.Error)
	return books
}

func (b *BookrespositoryImpl) FindbyId(booksId int) (model.Book, error) {
	var books model.Book
	result := b.Db.Find(&books, booksId)

	if result == nil {
		return books, errors.New("Book is not found")
	}

	return books, nil

}

func (b *BookrespositoryImpl) Save(book model.Book) {
	result := b.Db.Create(&book)
	helper.PanicError(result.Error)
}

func (b *BookrespositoryImpl) Update(book model.Book) {

	var updateag = request.UpdateBookRequest{
		Id:   book.Id,
		Name: book.Name,
	}

	result := b.Db.Model(&book).Updates(updateag)
	helper.PanicError(result.Error)
}
