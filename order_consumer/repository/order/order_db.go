package repository_order

import (
	"gorm.io/gorm"
)



type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	db.Table("orders").AutoMigrate(&Order{})
	db.Table("order_details").AutoMigrate(&Order_detail{})
	return orderRepository{db: db}
}

type OrderRepository interface {
	Create_order(order Order) error 
	Create_order_detail(order_detail Order_detail) error
}