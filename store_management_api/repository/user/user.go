package repository_user

import (
	"gorm.io/gorm"
	"time"
)


type userRepository struct {
	db *gorm.DB
}

type User struct {
	ID         string `gorm:"primary_key;not_null"`
	Username   string
	Password   string
	Email      string
	First_name string
	Last_name  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}


func NewUserRepository (db *gorm.DB) UserRepository {
	db.Table("users").AutoMigrate(&User{})
	return  userRepository{db}
}


type UserRepository interface {
	Create(user User) error
	Update(user User) error
	Delete(id string) error
	FindAll() (users []User, err error)
	FindbyId(id string) (user User, err error)
}