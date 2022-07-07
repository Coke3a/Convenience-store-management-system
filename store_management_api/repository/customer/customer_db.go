package repository_customer

func (db customerRepository) Create(customer Customer) error {
	err := db.db.Table("customers").Create(&customer).Error
	if err != nil {
		return err
	}
	return nil
}

func (db customerRepository) Update(customer Customer) error {
	err := db.db.Table("customers").Model(&Customer{}).Where("id=?", customer.ID).Updates(customer).Error
	if err != nil {
		return err
	}
	return nil
}

func (db customerRepository) Delete(id string) error {
	err := db.db.Table("customers").Where("id=?", id).Delete(&Customer{}).Error
	if err != nil {
		return err
	}
	return nil
}
func (db customerRepository) FindAll() (customers []Customer, err error) {
	err = db.db.Table("customers").Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}
func (db customerRepository) FindbyId(id string) (customer Customer, err error) {
	err = db.db.Table("customers").Where("id=?", id).First(&customer).Error
	if err != nil {
		return
	}
	return customer, nil
}
