package service_customer

import repository "github.com/Coke3a/Convenience-store-management-system/store_management/repository/customer"

type customerService struct {
	customerRepo repository.CustomerRepository
}

type CustomerResponse struct {
	ID               string `json:"id"`
	First_name       string `json:"first_name"`
	Last_name        string `json:"last_name"`
	Telephone_number string `json:"telephone_number"`
	Point            int    `json:"point"`
}

type CustomerService interface {
	AddNewCustomer(first_name string, last_name string, telephone_number string) error
	UpdateCustomer(id string, first_name string, last_name string, telephone_number string) error
	DeleteCustomer(id string) error
	FindAllCustomer() (customers []CustomerResponse, err error)
	FindCustomerById(id string) (customer CustomerResponse, err error)
}

func NewCustomerService(customerRepo repository.CustomerRepository) CustomerService {
	return customerService{customerRepo: customerRepo}
}
