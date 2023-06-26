package main

import (
	"github.com/Cekretik/BookApp/cmd/main/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.BookAppRoutes()
	r.GET("/route")
	r.Run("localhost:5432")
}
