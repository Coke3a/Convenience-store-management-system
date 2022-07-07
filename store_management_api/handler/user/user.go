package handler_user

import service "github.com/Coke3a/Convenience-store-management-system/store_management/services/user"



type userHandler struct {
	userServ service.UserService
}



type UserRequest struct{
	ID         string `json:"id"`
	Username   string `json:"user_name"`
	Password   string `json:"pass_word"`
	Email      string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}

func NewUserHandler(userServ service.UserService) userHandler {
	return userHandler{userServ: userServ}
}
