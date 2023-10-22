package routes

import (
	"github.com/Cekretik/BookApp/cmd/main/pkg/repositories"
	"github.com/gin-gonic/gin"
)

func BookAppRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/book", repositories.CreateBook)
	router.GET("/book", repositories.GetBook)
	router.GET("/book/:bookId", repositories.GetBookById)
	router.PUT("/book/:bookId", repositories.UpdateBook)
	router.DELETE("/book/:bookId", repositories.DeleteBook)
	return router
}
