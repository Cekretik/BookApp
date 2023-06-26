package routers

import (
	"github.com/Cekretik/BookApp/main/pkg/handlers"
	"github.com/gin-gonic/gin"
)

var BookAppRoutes = func() *gin.Engine {
	router := gin.Default()

	router.POST("/book", handlers.CreateBookHandler)
	router.GET("/book", handlers.GetBookHandler)
	router.GET("/book{bookId}", handlers.GetBookByIdHandler)
	router.PUT("/book{bookId}", handlers.UpdateBookHandler)
	router.DELETE("/book{bookId}", handlers.DeleteBookHandler)
	return router
}
