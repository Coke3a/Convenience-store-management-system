package commands

import "time"

type Order_detail_Command struct {
	ID              string `json:"id"`
	Product_id      string `json:"product_id"`
	Order_id        string `json:"order_id"`
	Product_barcode string `json:"product_barcode"`
	Product_name    string `json:"product_name"`
	Product_Price   int    `json:"product_price"`
	CreatedAt       time.Time
}

type Order_Command struct {
	ID                  string `json:"id"`
	Customer_ID         string `json:"customer_id"`
	Customer_point      int    `json:"customer_point"`
	Promotion_ID        string `json:"promotion_id"`
	Use_point           int    `json:"use_point"`
	Total_without_Promo int    `json:"total_without_promo"`
	Discount            int    `json:"discount"`
	Total_Price         int    `json:"total_price"`
	CreatedAt           time.Time
}

type Product_Command struct {
	ID              string `json:"id"`
	Product_barcode string `json:"product_barcode"`
	Name            string `json:"name"`
	Price           int    `json:"price"`
}

type Customer_command struct {
	ID               string `json:"id"`
	First_name       string `json:"first_name"`
	Last_name        string `json:"last_name"`
	Telephone_number string `json:"telephone_number"`
	Point            int    `json:"point"`
}

type Promotion_Command struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	PurchaseMin    int    `json:"purchase_min"`
	Required_point int    `json:"required_point"`
	Type           int    `json:"type"`
	Discount       int    `json:"discount"`
}
