package main

import (
	"bookapi/config"
	"bookapi/controller"
	"bookapi/helper"
	"bookapi/model"
	"bookapi/repository"
	"bookapi/routers"
	"net/http"

	"bookapi/service"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/rs/cors"
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

	// Enable CORS using the default options. You can customize it further as needed.
	corsHandler := cors.Default().Handler(router)

	server := &http.Server{
		Addr:    ":3000",
		Handler: corsHandler, // Use the corsHandler as the handler for the server.
	}

	err := server.ListenAndServe()

	helper.PanicError(err)
}
