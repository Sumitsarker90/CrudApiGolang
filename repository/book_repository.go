package repository

import (
	"bookapi/model"
)

type Bookrespository interface {
	Save(books model.Book)
	Update(books model.Book)
	Delete(booksId int)
	FindbyId(booksId int) (book model.Book, err error)
	Findall() []model.Book
}
