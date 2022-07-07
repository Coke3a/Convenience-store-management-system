package repository_customer

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	ID               string `gorm:"primary_key;not_null"`
	First_name       string
	Last_name        string
	Telephone_number string
	Point            int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type customerRepository struct {
	db *gorm.DB
}

type CustomerRepository interface {
	Create(customer Customer) error
	Update(customer Customer) error
	Delete(id string) error
	FindAll() (customers []Customer, err error)
	FindbyId(id string) (customer Customer, err error)
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	db.Table("customers").AutoMigrate(&Customer{})
	return customerRepository{db}
}
