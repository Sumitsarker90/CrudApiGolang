package service

import (
	request "bookapi/data/request/response"
)

type BooksService interface {
	Create(books request.CreateBookRequest)
	Update(books request.UpdateBookRequest)
	Delete(booksId int)
	FindById(booksId int) request.BooksResponse
	FindAll() []request.BooksResponse
}
