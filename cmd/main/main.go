package main

import (
	"github.com/Cekretik/BookApp/cmd/main/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.BookAppRoutes().Run(":8080")
	routes.UserRoutes().Run(":8080")
	router.Run(":8080")
}
