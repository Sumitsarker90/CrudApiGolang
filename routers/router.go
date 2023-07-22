package routers

import (
	"bookapi/controller"
	

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(booksController *controller.BooksController) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	booksRouter := baseRouter.Group("/books")
	booksRouter.GET("", booksController.FindAll)
	booksRouter.GET("/:bookId", booksController. FindByid)
	booksRouter.POST("", booksController.Create)
	booksRouter.PATCH("/:bookId", booksController.Update)
	booksRouter.DELETE("/:bookId", booksController.Delete)

	return router
}
