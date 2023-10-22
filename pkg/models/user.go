package models

import (
	"github.com/Cekretik/BookApp/cmd/main/pkg/config"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username;unique_index" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Email    string `gorm:"column:email;unique_index" json:"email"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func CreateUserFromDB(user *User) (*User, error) {
	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func GetUserByIDFromDB(id uint) (*User, error) {
	var user User
	result := db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetAllUsersFromDB() ([]User, error) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func UpdateUserFromDB(ID uint, userToUpdate *User) error {
	var existingUser User
	result := db.First(&existingUser, ID)
	if result.Error != nil {
		return result.Error
	}

	if userToUpdate.Username != "" {
		existingUser.Username = userToUpdate.Username
	}
	if userToUpdate.Password != "" {
		existingUser.Password = userToUpdate.Password
	}
	if userToUpdate.Email != "" {
		existingUser.Email = userToUpdate.Email
	}

	result = db.Save(&existingUser)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteUserFromDB(ID uint) error {
	result := db.Delete(&User{}, ID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
