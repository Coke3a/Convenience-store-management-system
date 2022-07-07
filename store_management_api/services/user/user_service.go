package service_user

import (
	"log"

	repository "github.com/Coke3a/Convenience-store-management-system/store_management/repository/user"
	"github.com/google/uuid"
)


func (r userService) AddNewUser(username string, password string, email string, first_name string, last_name string) error{
	user := repository.User{
		ID: uuid.NewString(),
		Username: username,
		Password: password,
		Email: email,
		First_name: first_name,
		Last_name: last_name,
	 }
	err := r.userRepo.Create(user)
	if err != nil {
	   log.Println(err)
	   return err 
   }
	   return nil 
   }
   
   func (r userService) UpdateUser(id string , username string, password string, email string, first_name string, last_name string) error{
	user := repository.User{
		ID: id,
		Username: username,
		Password: password,
		Email: email,
		First_name: first_name,
		Last_name: last_name,
	 }
		err := r.userRepo.Update(user)
		if err != nil {
		   log.Println(err)
		   return err 
	   }
	   return nil 
   
   }
   func (r userService) DeleteUser(id string) error {
	err := r.userRepo.Delete(id)
	if err != nil {
	   log.Println(err)
	   return err 
   }
   return nil

   }
   
   func (r userService) FindAllUsers()(users []UserResponse, err error){
	   userDB, err := r.userRepo.FindAll()
	   if err != nil {
		   log.Println(err)
		   return nil, err 
	   }
	   for _, u := range userDB {
		   users = append(users, UserResponse{
			   ID: u.ID,
			   Username: u.Username,
			   Email: u.Email,
			   First_name: u.First_name,
			   Last_name: u.Last_name,
		   })
	   }
	   return users , nil 
   }
   
   func (r userService) FindUserByID(id string) (user UserResponse, err error){
	   userDB, err := r.userRepo.FindbyId(id)
	   if err != nil {
		   log.Println(err)
		   return 
	   }
	   user = UserResponse{
		   ID: userDB.ID,
		   Username: userDB.Username,
		   Email: userDB.Email,
		   First_name: userDB.First_name,
		   Last_name: userDB.Last_name,
	   }
	   return user , nil 
   }