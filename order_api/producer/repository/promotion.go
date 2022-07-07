package repository_producer

import "github.com/Coke3a/Convenience-store-management-system/order/producer/commands"

func (db orderRepository) GetPromotionwithConditions(total int, point int) (promotions []commands.Promotion_Command, err error) {
	err = db.db.Table("promotions").Where("purchase_min <= ? AND required_point <= 0", total).Find(&promotions).Error
	if err != nil {
		return nil, err
	}
	return promotions, nil
}



func (db orderRepository) GetPromotionbyId(id string) (promotion commands.Promotion_Command, err error) {
	err = db.db.Table("promotions").Where("id=?",id).First(&promotion).Error
	if err != nil {
		return   
	}
	return promotion, nil
}
