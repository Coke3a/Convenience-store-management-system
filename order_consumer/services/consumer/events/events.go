package handler_consumer

import "reflect"

var Topics = []string{
	reflect.TypeOf(SubmitOrderEvent{}).Name(),
	reflect.TypeOf(SubmitOrder_Order_DetailEvent{}).Name(),
}


type SubmitOrderEvent struct {
	ID                  string `json:"id"`
	Customer_ID         string `json:"customer_id"`
	Customer_point      int    `json:"customer_point"`
	Promotion_ID        string `json:"promotion_id"`
	Use_point           int    `json:"use_point"`
	Total_without_Promo int    `json:"total_without_promo"`
	Discount            int    `json:"discount"`
	Total_Price         int    `json:"total_price"`
}

type SubmitOrder_Order_DetailEvent struct {
	ID              string `json:"id"`
	Product_id      string `json:"product_id"`
	Order_id        string `json:"order_id"`
	Product_barcode string `json:"product_barcode"`
	Product_name    string `json:"product_name"`
	Product_Price   int    `json:"product_price"`
}
