package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type User struct {
	ID       uint   `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Email    string `gorm:"column:email" json:"email"`
}

func (*User) Error() string {
	panic("unimplemented")
}

var db *gorm.DB

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(c *gin.Context) {
	var creds Credentials

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid json provided", "details": err.Error()})
		return
	}

	user, err := getUserByUsername(creds.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found", "details": err.Error()})
		return
	}

	if !checkPasswordHash(creds.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password", "details": err.Error()})
		return
	}

	token, err := verifyToken(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to generate token", "details": err.Error()})
		return
	}
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
}

func Register(c *gin.Context) {
	var creds Credentials

	newUser := User{
		Username: creds.Username,
		Password: creds.Password,
		Email:    creds.Email,
	}

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid json provided", "details": err.Error()})
		return
	}

	existingUser, err := getUserByUsername(creds.Username)
	if err != nil && existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with the same username already exists", "details": err.Error()})
		return

	}

	existingEmail, err := getUserByUsername(creds.Email)
	if err != nil && existingEmail != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with the same email already exists", "details": err.Error()})
		return
	}

	err = createUser(newUser.Username, newUser.Password, newUser.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user", "details": err.Error()})
		return
	}

	token, err := generateToken(newUser.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to generate token", "details": err.Error()})
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "You have successfully registered", "details": err.Error()})
}

func getUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func createUser(username string, password string, email string) error {
	var user User
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
