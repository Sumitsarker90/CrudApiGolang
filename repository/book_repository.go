package repository

import (
	"bookapi/model"
)

type Bookrespository interface {
	Save(book model.Book)
	Update(book model.Book)
	Delete(bookId int)
	FindbyId(bookId int) (book model.Book, err error)
	Findall() []model.Book
}
