package models

type User struct {
	ID       uint   `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Email    string `gorm:"column:email" json:"email"`
}

type APIResponse struct {
	Status  int
	Message string
	Data    interface{}
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

func UpdateUserFromDB(ID uint) error {
	result := db.Save(ID)
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

func CreateUser(user *User) APIResponse {
	user, err := CreateUserFromDB(user)
	if err != nil {
		return APIResponse{
			Status:  500,
			Message: "Failed to create user",
		}
	}
	return APIResponse{
		Status:  200,
		Message: "Success",
		Data:    user,
	}
}

func GetUserByID(id uint) APIResponse {
	user, err := GetUserByIDFromDB(id)
	if err != nil {
		return APIResponse{
			Status:  404,
			Message: "User not found",
		}
	}
	return APIResponse{
		Status:  200,
		Message: "Success",
		Data:    user,
	}
}

func GetAllUsers() APIResponse {
	users, err := GetAllUsersFromDB()
	if err != nil {
		return APIResponse{
			Status:  500,
			Message: "Failed to get all users",
		}
	}
	return APIResponse{
		Status:  200,
		Message: "Success",
		Data:    users,
	}
}

func UpdateUser(ID uint) APIResponse {
	err := UpdateUserFromDB(ID)
	if err != nil {
		return APIResponse{
			Status:  500,
			Message: "Failed to update user",
		}
	}
	return APIResponse{
		Status:  200,
		Message: "Success",
		Data:    ID,
	}
}

func DeleteUser(ID uint) APIResponse {
	err := DeleteUserFromDB(ID)
	if err != nil {
		return APIResponse{
			Status:  500,
			Message: "Failed to delete user",
		}
	}
	return APIResponse{
		Status:  200,
		Message: "Success",
	}
}
