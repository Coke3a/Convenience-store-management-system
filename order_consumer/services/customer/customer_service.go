package service_customer

import (
	"log"

	repository "github.com/Coke3a/Convenience-store-management-system/order_consumer/repository/customer"
)



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

func (r customerService) UsePoints(id string, newPoint int) error {
	customer, err := r.FindCustomerById(id)
	if err != nil {
		log.Println(err)
		return err
	}
		customer.Point = newPoint
		updateCustomer := repository.Customer{
			ID:               customer.ID,
			First_name:       customer.First_name,
			Last_name:        customer.Last_name,
			Telephone_number: customer.Telephone_number,
			Point:            customer.Point,
		}
		err = r.customerRepo.Update(updateCustomer)
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}


////////
// func (r customerService)	EarnPoint(id string, total_price int) error {

// }
