package service

import (
	request "bookapi/data/request/response"
	"bookapi/helper"
	"bookapi/model"
	"bookapi/repository"

	"github.com/go-playground/validator/v10"
)

type BooksServiceeImpl struct {
	BooksRepository repository.Bookrespository
	validate        *validator.Validate
}

func NewBoooksServiceImpl(bookRepository repository.Bookrespository, validate *validator.Validate) BooksService {
	return &BooksServiceeImpl{

		BooksRepository: bookRepository,
		validate:        validate,
	}
}


// Create implements BooksService.
func (b*BooksServiceeImpl) Create(books request.CreateBookRequest) {
	 err := b.validate.Struct(books)
	 helper.PanicError(err)
	 tagModel :=model.Book{
	 	
	 	Name: books.Name,
	 }
	 b.BooksRepository.Save(tagModel)


}

// Delete implements BooksService.
func ( b *BooksServiceeImpl) Delete(booksId int) {
	b.BooksRepository.Delete(booksId)
}

// FindAll implements BooksService.
func (b *BooksServiceeImpl) FindAll() []request.BooksResponse {
	result := b.BooksRepository.Findall()

	var books[]request.BooksResponse
	for _, value := range result{
		book := request.BooksResponse{
			Id: value.Id,
			Name: value.Name,
		}
		books= append(books, book)
	}
	  return books

}

// FindById implements BooksService.
func (b *BooksServiceeImpl) FindById(booksId int) request.BooksResponse {
	tagData, err := b.BooksRepository.FindbyId(booksId)
	helper.PanicError(err)

	tagResponse := request.BooksResponse{
		Id: tagData.Id,
		Name: tagData.Name,
	}

	return tagResponse
}

// Update implements BooksService.

func (b *BooksServiceeImpl) Update(request request.UpdateBookRequest) {
	tagData, err := b.BooksRepository.FindbyId(request.Id)
	helper.PanicError(err)
	tagData.Name = request.Name

	b.BooksRepository.Update(tagData)
}



