package repositories

import (
	"net/http"
	"strconv"

	"github.com/Cekretik/BookApp/cmd/main/pkg/models"
	"github.com/Cekretik/BookApp/cmd/main/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var requestUser models.User
	utils.ParseBody(c, &requestUser)
	response := models.CreateUser(&requestUser)
	if response.Status != 200 {
		c.JSON(response.Status, gin.H{"error": "Failed to create user", "details": response.Message})
		return
	}

	respondWithJSON(c, http.StatusOK, requestUser)
}

func GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	ID, err := strconv.Atoi(userID)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid ID", err)
		return
	}

	userByID, err := models.GetUserByIDFromDB(uint(ID))
	if err != nil {
		respondWithError(c, http.StatusNotFound, "User not found", err)
		return
	}

	respondWithJSON(c, http.StatusOK, userByID)
}

func GetAllUsers(c *gin.Context) {
	allUsers, err := models.GetAllUsersFromDB()
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to get all users", err)
		return
	}

	respondWithJSON(c, http.StatusOK, allUsers)
}

func UpdateUser(c *gin.Context) {
	var requestUser models.User
	utils.ParseBody(c, &requestUser)

	userID := c.Param("id")
	ID, err := strconv.Atoi(userID)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid ID", err)
		return
	}

	err = models.UpdateUserFromDB(uint(ID), &requestUser)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to update user", err)
		return
	}

	respondWithMessage(c, http.StatusOK, "User updated successfully")
}

func DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	ID, err := strconv.Atoi(userID)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid ID", err)
		return
	}

	err = models.DeleteUserFromDB(uint(ID))
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to delete user", err)
		return
	}

	respondWithMessage(c, http.StatusOK, "User deleted successfully")
}

func respondWithError(c *gin.Context, httpStatus int, message string, err error) {
	c.JSON(httpStatus, gin.H{"error": message, "details": err.Error()})
}

func respondWithJSON(c *gin.Context, httpStatus int, data interface{}) {
	c.JSON(httpStatus, data)
}

func respondWithMessage(c *gin.Context, httpStatus int, message string) {
	c.JSON(httpStatus, gin.H{"message": message})
}
