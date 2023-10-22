package routes

import (
	"net/http"
	"strconv"

	"github.com/Cekretik/BookApp/cmd/main/pkg/models"
	"github.com/Cekretik/BookApp/cmd/main/pkg/repositories"
	"github.com/Cekretik/BookApp/cmd/main/pkg/utils"
	"github.com/gin-gonic/gin"
)

func UserRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/users", repositories.CreateUser, func(c *gin.Context) {
		user := models.User{}
		utils.ParseBody(c, &user)
		response := models.CreateUser(&user)
		if response.Status != 200 {
			c.JSON(response.Status, gin.H{"error": "Failed to create user", "details": response.Message})
			return
		}
		c.JSON(response.Status, response)
	})

	router.GET("/users/:id", repositories.GetUserByID, func(c *gin.Context) {
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

	router.GET("/users", repositories.GetAllUsers, func(c *gin.Context) {
		response := models.GetAllUsers()
		if response.Status != 200 {
			c.JSON(response.Status, gin.H{"error": "Failed to get all users", "details": response.Message})
			return
		}
		c.JSON(response.Status, response)
	})

	router.DELETE("/users/:id", repositories.DeleteUser, func(c *gin.Context) {
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

	router.PUT("/users/:id", repositories.UpdateUser, func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil || id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID", "details": "Invalid or missing user ID"})
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
