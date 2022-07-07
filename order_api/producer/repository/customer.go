package repository_producer

import "github.com/Coke3a/Convenience-store-management-system/order/producer/commands"



func (db orderRepository) GetCustomerByTelephone_number(telephone_number string) (customer commands.Customer_command, err error) {
	err = db.db.Table("customers").Where("telephone_number=?", telephone_number).First(&customer).Error
	if err != nil {
	return customer, err  
	}
	return customer, nil
}
