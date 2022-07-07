package handler_customer

import service "github.com/Coke3a/Convenience-store-management-system/store_management/services/customer"


type CustomerRequest struct {
	ID               string `json:"id"`
	First_name       string `json:"first_name"`
	Last_name        string `json:"last_name"`
	Telephone_number string `json:"telephone_number"`
	Point            int    `json:"point"`
}

type customerHandler struct {
	customerSev service.CustomerService
}

func NewCustomerHandler(customerSev service.CustomerService) customerHandler {
	return customerHandler{customerSev: customerSev}
}
