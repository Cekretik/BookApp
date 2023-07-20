package main

import (
	"github.com/Cekretik/BookApp/cmd/main/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.BookAppRoutes()
	router.GET("/route", func(c *gin.Context) {})

	router.Run("localhost:5432")
}
