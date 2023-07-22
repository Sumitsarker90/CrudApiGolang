package controller

import (
	request "bookapi/data/request/response"
	"bookapi/helper"
	"bookapi/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type BooksController struct {
	booksService service.BooksService
}

func NewBooksController(service service.BooksService) *BooksController {
	return &BooksController{
		booksService: service,
	}
}

func (controller *BooksController) Create(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	createBookRequest := request.CreateBookRequest{}
	err := ctx.ShouldBindJSON(&createBookRequest)
	helper.PanicError(err)

	controller.booksService.Create(createBookRequest)
	webResponse := request.WebResponse{ // Can be an error in future
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("content-Type", "application/json")

	ctx.JSON(http.StatusOK, webResponse)
}

// Update controller
func (controller *BooksController) Update(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	updateBookrequest := request.UpdateBookRequest{}
	err := ctx.ShouldBindJSON(&updateBookrequest)
	helper.PanicError(err)

	bookId := ctx.Param("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicError(err)
	updateBookrequest.Id = id

	controller.booksService.Update(updateBookrequest)

	webResponse := request.WebResponse{ // Can be an error in future
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// Delete controller
func (controller *BooksController) Delete(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	bookId := ctx.Param("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicError(err)
	controller.booksService.Delete(id)

	webResponse := request.WebResponse{ // Can be an error in future
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// Findbyid Controller
func (controller *BooksController) FindByid(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	bookId := ctx.Param("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicError(err)
	tagResponse := controller.booksService.FindById(id)

	webResponse := request.WebResponse{ // Can be an error in future
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll Controller

func (controller *BooksController) FindAll(ctx *gin.Context) {

	log.Info().Msg("findAll tags")
	tagResponse := controller.booksService.FindAll()
	webResponse := request.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
