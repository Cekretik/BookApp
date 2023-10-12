package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var creds Credentials

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := getUserByUsername(creds.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !user.checkPasswordHash(creds.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}
}

func Register(c *gin.Context) {
	var creds Credentials

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "You have successfully registered"})

}

func getUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
