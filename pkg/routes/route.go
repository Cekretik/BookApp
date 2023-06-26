package routes

import (
	"github.com/Cekretik/BookApp/cmd/main/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var BookAppRoutes = func() *gin.Engine {
	router := gin.Default()

	router.POST("/book", controllers.CreateBook)
	router.GET("/book", controllers.GetBook)
	router.GET("/book{bookId}", controllers.GetBookById)
	router.PUT("/book{bookId}", controllers.UpdateBook)
	router.DELETE("/book{bookId}", controllers.DeleteBook)
	return router
}
