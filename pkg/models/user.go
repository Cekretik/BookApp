package models

type User struct {
	ID       uint   `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Email    string `gorm:"column:email" json:"email"`
}

func CreateUser(user *User) (*User, error) {
	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func GetUserById(id uint) (*User, error) {
	var user User
	result := db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func UpdateUser(user *User) (*User, error) {
	result := db.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func DeleteUser(ID uint) error {
	result := db.Delete(&User{}, ID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
