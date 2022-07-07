package service_customer

import repository "github.com/Coke3a/Convenience-store-management-system/order_consumer/repository/customer"

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
	FindCustomerById(id string) (customer CustomerResponse, err error)
	UsePoints(id string, newpoint int) error
}

func NewCustomerService(customerRepo repository.CustomerRepository) CustomerService {
	return customerService{customerRepo: customerRepo}
}
