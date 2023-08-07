package routers

import (
	"bookapi/controller"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(booksController *controller.BooksController) *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 21 //16 MB Max file size
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowMethods("GET", "POST", "DELETE", "PUT", "OPTIONS", "HEAD")
	corsConfig.AddAllowHeaders("Authorization", "WWW-Authenticate", "Content-Type", "Accept", "X-Requested-With")
	corsConfig.AddExposeHeaders("Authorization", "WWW-Authenticate", "Content-Type", "Accept", "X-Requested-With")
	router.Use(cors.New(corsConfig))
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	booksRouter := baseRouter.Group("/books")
	booksRouter.GET("", booksController.FindAll)
	booksRouter.GET("/:bookId", booksController.FindByid)
	booksRouter.POST("/insert", booksController.Create)
	booksRouter.PATCH("/:bookId", booksController.Update)
	booksRouter.DELETE("/:bookId", booksController.Delete)

	return router
}
