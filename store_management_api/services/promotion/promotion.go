package service_promotion

import repository "github.com/Coke3a/Convenience-store-management-system/store_management/repository/promotion"

type promotionService struct {
	promotionRepo repository.PromotionRepository
}

type PromotionResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Type           int    `json:"type"`
	PurchaseMin    int    `json:"purchase_min"`
	Required_point int    `json:"required_point"`
	Discount       int    `json:"discount"`
}

type PromotionService interface {
	// AddNewPromotion(name string, purchaseMin uint, required_point uint, discountPrice uint, discountPercent uint) error
	AddNewPromotionDiscountPrice(name string, purchaseMin int, required_point int, discount int) error
	AddNewPromotionDiscountPercent(name string, purchaseMin int, required_point int, discount int) error
	UpdatePromotion(id string, name string, purchaseMin int, required_point int, discount int) error
	DeletePromotion(id string) error
	FindAllPromotion() (promotions []PromotionResponse, err error)
	FindPromotionById(id string) (promotion PromotionResponse, err error)
}

func NewPromotionService(promotionRepo repository.PromotionRepository) PromotionService {
	return promotionService{promotionRepo: promotionRepo}
}
