package repository_product

func (db productRepository) Create(product Product) error {
	err := db.db.Table("products").Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (db productRepository) Update(product Product) error {
	err := db.db.Table("products").Model(&Product{}).Where("id=?", product.ID).Updates(product).Error
	if err != nil {
		return err
	}
	return nil
}

func (db productRepository) Delete(id string) error {
	err := db.db.Table("products").Where("id=?", id).Delete(&Product{}).Error
	if err != nil {
		return err
	}
	return nil

}
func (db productRepository) FindAll() (products []Product, err error) {
	err = db.db.Table("products").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (db productRepository) FindbyId(id string) (product Product, err error) {
	err = db.db.Table("products").Where("id=?", id).First(&product).Error
	if err != nil {
		return
	}
	return product, nil
}
