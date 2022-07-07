package repository_promotion

func (db promotionRepository) Create(promotion Promotion) error {
	err := db.db.Table("promotions").Create(&promotion).Error
	if err != nil {
		return err
	}
	return nil
}

func (db promotionRepository) Update(promotion Promotion) error {
	err := db.db.Table("promotions").Model(&Promotion{}).Where("id=?", promotion.ID).Updates(promotion).Error
	if err != nil {
		return err
	}
	return nil
}

func (db promotionRepository) Delete(id string) error {
	err := db.db.Table("promotions").Where("id=?", id).Delete(&Promotion{}).Error
	if err != nil {
		return err
	}
	return nil

}
func (db promotionRepository) FindAll() (promotions []Promotion, err error) {
	err = db.db.Table("promotions").Find(&promotions).Error
	if err != nil {
		return nil, err
	}
	return promotions, nil
}
func (db promotionRepository) FindbyId(id string) (promotion Promotion, err error) {
	err = db.db.Table("promotions").Where("id=?", id).First(&promotion).Error
	if err != nil {
		return
	}
	return promotion, nil
}
