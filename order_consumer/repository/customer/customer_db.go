package repository_customer


func (db customerRepository) FindbyId(id string) (customer Customer, err error) {
	err = db.db.Table("customers").Where("id=?", id).First(&customer).Error
	if err != nil {
		return
	}
	return customer, nil
}


func (db customerRepository) Update(customer Customer) error {
	err := db.db.Table("customers").Model(&Customer{}).Where("id=?", customer.ID).Updates(customer).Error
	if err != nil {
		return err
	}
	return nil
}