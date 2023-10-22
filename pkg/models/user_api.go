package models

type APIResponse struct {
	Status  int
	Message string
	Data    interface{}
}

func CreateUser(user *User) APIResponse {
	createdUser, err := CreateUserFromDB(user)
	if err != nil {
		return APIResponse{
			Status:  500,
			Message: "Failed to create user",
		}
	}
	return APIResponse{
		Status:  200,
		Message: "Success",
		Data:    createdUser,
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
	err := UpdateUserFromDB(ID, &User{})
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
