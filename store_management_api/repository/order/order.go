package repository_order

import "time"
type Order_detail struct {
	ID              string `gorm:"primary_key;not_null"`
	Product_id      string
	Order_id        string
	Product_barcode string
	Product_name    string
	Product_Price   int
	CreatedAt       time.Time
}

type Order struct {
	ID                  string `gorm:"primary_key;not_null"`
	Customer_ID         string
	Customer_point      int
	Promotion_ID        string
	Use_point           int
	Total_without_Promo int
	Discount            int
	Total_Price         int
	CreatedAt           time.Time
}