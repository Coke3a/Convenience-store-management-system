package service_promotion

import (
	"log"

	repository "github.com/Coke3a/Convenience-store-management-system/store_management/repository/promotion"
	"github.com/google/uuid"
)





func (r promotionService)AddNewPromotionDiscountPrice(name string, purchaseMin int, required_point int, discount int) error{
	
	promotion := repository.Promotion{
		ID: uuid.NewString(),
		Name: name,
		PurchaseMin: purchaseMin,
		Required_point: required_point,
		Type: 1,
		Discount: discount,
	}
	err := r.promotionRepo.Create(promotion)
	if err != nil {
		log.Println(err)
		return err 
	}
	return nil 
}

func (r promotionService)AddNewPromotionDiscountPercent(name string, purchaseMin int, required_point int, discount int) error{
	
	
	promotion := repository.Promotion{
		ID: uuid.NewString(),
		Name: name,
		PurchaseMin: purchaseMin,
		Required_point: required_point,
		Type: 2,
		Discount: discount,
	}
	err := r.promotionRepo.Create(promotion)
	if err != nil {
		log.Println(err)
		return err 
	}
	return nil 
}


func (r promotionService)UpdatePromotion(id string, name string, purchaseMin int, required_point int, discount int) error{
	promotion := repository.Promotion{
		ID: id,
		Name: name,
		PurchaseMin: purchaseMin,
		Required_point: required_point,
		Discount: discount,
	}

	err := r.promotionRepo.Update(promotion)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil 
}




func (r promotionService)DeletePromotion(id string) error {
	err := r.promotionRepo.Delete(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}





func (r promotionService)FindAllPromotion() (promotions []PromotionResponse , err error) {
	promotionDB, err  := r.promotionRepo.FindAll()
	if err != nil {
		log.Println(err)
		return nil, err 
	}
	for _, p := range promotionDB {
		promotions = append(promotions, PromotionResponse{
			ID: p.ID,
			Name: p.Name,
			Type: p.Type,
			PurchaseMin: p.PurchaseMin,
			Required_point: p.Required_point,
			Discount: p.Discount,
		})
	}
	return promotions, nil 
}

func (r promotionService)FindPromotionById(id string) (promotion PromotionResponse, err error)  {
	promotionDB, err := r.promotionRepo.FindbyId(id)
	if err != nil {
		log.Println(err)
		return 
	}

	promotion = PromotionResponse{
		ID: promotionDB.ID,
		Name: promotionDB.Name,
		PurchaseMin: promotionDB.PurchaseMin,
		Required_point: promotionDB.Required_point,
		Discount: promotionDB.Discount,
	}
	return promotion, nil 

}