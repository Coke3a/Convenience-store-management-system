package handler_promotion

import service "github.com/Coke3a/Convenience-store-management-system/store_management/services/promotion"


type PromotionRequest struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Type           int    `json:"type"`
	PurchaseMin    int    `json:"purchase_min"`
	Required_point int    `json:"required_point"`
	Discount       int    `json:"discount"`
}


type promotionHandler struct {
	promotionServ service.PromotionService
}

func NewPromotionHandler(promotionServ service.PromotionService) promotionHandler {
	return promotionHandler{promotionServ: promotionServ}
}