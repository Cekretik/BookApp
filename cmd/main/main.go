package main

import (
	"github.com/Cekretik/BookApp/cmd/main/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	go routes.BookAppRoutes().Run(":8080")
	go routes.UserRoutes().Run(":8081")

	router.Run(":8082")
}
