package service_customer

import (
	"log"

	repository "github.com/Coke3a/Convenience-store-management-system/store_management/repository/customer"
	"github.com/google/uuid"
)

func (r customerService) AddNewCustomer(first_name string, last_name string, telephone_number string) error {
	customer := repository.Customer{
		ID:               uuid.NewString(),
		First_name:       first_name,
		Last_name:        last_name,
		Telephone_number: telephone_number,
	}
	err := r.customerRepo.Create(customer)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (r customerService) UpdateCustomer(id string, first_name string, last_name string, telephone_number string) error {
	customer := repository.Customer{
		ID:               id,
		First_name:       first_name,
		Last_name:        last_name,
		Telephone_number: telephone_number,
	}
	err := r.customerRepo.Update(customer)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r customerService) DeleteCustomer(id string) error {
	err := r.customerRepo.Delete(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r customerService) FindAllCustomer() (customers []CustomerResponse, err error) {
	customersDB, err := r.customerRepo.FindAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, p := range customersDB {
		customers = append(customers, CustomerResponse{
			ID:               p.ID,
			First_name:       p.First_name,
			Last_name:        p.Last_name,
			Telephone_number: p.Telephone_number,
			Point:            p.Point,
		})
	}
	return customers, nil
}

func (r customerService) FindCustomerById(id string) (customer CustomerResponse, err error) {
	customerDB, err := r.customerRepo.FindbyId(id)
	if err != nil {
		log.Println(err)
		return
	}
	customer = CustomerResponse{
		ID:               customerDB.ID,
		First_name:       customerDB.First_name,
		Last_name:        customerDB.Last_name,
		Telephone_number: customerDB.Telephone_number,
		Point:            customerDB.Point,
	}
	return customer, nil
}
