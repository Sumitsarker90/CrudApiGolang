package main

import (
	"bookapi/config"
	"bookapi/controller"
	"bookapi/helper"
	"bookapi/model"
	"bookapi/repository"
	"bookapi/routers"
	
	"bookapi/service"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("Go Api")

	log.Info().Msg("Server started")

	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("books").AutoMigrate(&model.Book{})

	// Repository

	bookResository := repository.NewBookrepo(db)

	// service

	booksService := service.NewBoooksServiceImpl(bookResository, validate)

	// controllers
// Assuming you have imported the necessary packages correctly.

// Create a new BooksServiceImpl using booksService.
booksController := controller.NewBooksController(booksService)

// Create the router using the booksController as an input.
router := routers.NewRouter(booksController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}

	err := server.ListenAndServe()

	helper.PanicError(err)
}
