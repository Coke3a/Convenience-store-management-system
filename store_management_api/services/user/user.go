package service_user

import repository "github.com/Coke3a/Convenience-store-management-system/store_management/repository/user"



type userService struct {
	userRepo repository.UserRepository
}

type UserResponse struct{
	ID         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}


type UserService interface {
	AddNewUser(username string, password string, email string, first_name string, last_name string) error
	UpdateUser(id string, username string, password string, email string, first_name string, last_name string) error
	DeleteUser(id string) error
	FindAllUsers()(users []UserResponse, err error)
	FindUserByID(id string) (user UserResponse, err error)
}

func NewUserService (userRepo repository.UserRepository) UserService {
return userService{userRepo: userRepo}
}