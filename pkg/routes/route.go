package routes

import (
	"net/http"
	"strconv"

	"github.com/Cekretik/BookApp/cmd/main/pkg/models"
	"github.com/Cekretik/BookApp/cmd/main/pkg/repositories"
	"github.com/gin-gonic/gin"
)

var BookAppRoutes = func() *gin.Engine {
	router := gin.Default()

	router.POST("/book", repositories.CreateBook)
	router.GET("/book", repositories.GetBook)
	router.GET("/book{bookId}", repositories.GetBookById)
	router.PUT("/book{bookId}", repositories.UpdateBook)
	router.DELETE("/book{bookId}", repositories.DeleteBook)
	return router
}

var userRoutes = func() *gin.Engine {
	router := gin.Default()

	router.POST("/users", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided", "details": err.Error()})
			return
		}
		response := models.CreateUser(&user)
		if response.Status != 200 {
			c.JSON(response.Status, gin.H{"error": "Failed to create user", "details": response.Message})
			return
		}
		c.JSON(response.Status, response)
	})

	router.GET("/users/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID", "details": err.Error()})
			return
		}
		response := models.GetUserByID(uint(id))
		if response.Status != 200 {
			c.JSON(response.Status, gin.H{"error": "User not found", "details": response.Message})
			return
		}
		c.JSON(response.Status, response)
	})

	router.GET("/users", func(c *gin.Context) {
		response := models.GetAllUsers()
		if response.Status != 200 {
			c.JSON(response.Status, gin.H{"error": "Failed to get all users", "details": response.Message})
			return
		}
		c.JSON(response.Status, response)

	})

	router.GET("/users/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID", "details": err.Error()})
			return
		}

		response := models.GetUserByID(uint(id))
		if response.Status != 200 || response.Data == nil {
			c.JSON(response.Status, gin.H{"error": "User not found", "details": response.Message})
			return
		}
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Delete user failed", "details": err.Error()})
			return
		}
		response := models.DeleteUser(uint(id))
		if response.Status != 200 {
			c.JSON(response.Status, gin.H{"error": "Failed to delete user", "details": response.Message})
			return
		}
		c.JSON(response.Status, response)
	})

	router.PUT("/users/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID", "details": err.Error()})
			return
		}

		var userToUpdate models.User
		if err := c.ShouldBindJSON(&userToUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON provided", "details": err.Error()})
			return
		}

		response := models.UpdateUser(uint(id))
		if response.Status != 200 {
			c.JSON(response.Status, gin.H{"error": "Failed to update user", "details": response.Message})
			return
		}

		c.JSON(response.Status, gin.H{"message": "User updated successfully"})
	})

	return router
}
