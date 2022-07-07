package repository_producer

import (
	"github.com/Coke3a/Convenience-store-management-system/order/producer/commands"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	db.Table("tempdb_orders").AutoMigrate(&commands.Order_Command{})
	db.Table("tempdb_order_details").AutoMigrate(&commands.Order_detail_Command{})
	return orderRepository{db: db}
}

type OrderRepository interface {
	CreateOrder(id string) (err error)
	UpdateOrder(id string, order commands.Order_Command) error
	GetOrderById(id string) (order commands.Order_Command, err error)
	GetLastOrder() (order commands.Order_Command, err error)
	CreateOrderdetail(order_detail commands.Order_detail_Command) (err error)
	GetAllOrder_DetailByOrderId(order_id string) (order_details []commands.Order_detail_Command, err error)
	DeleteOrderDetailByID(id string) error
	GetProductById(product_code string) (product commands.Product_Command, err error)
	GetPromotionwithConditions(total int, point int) (promotions []commands.Promotion_Command, err error)
	GetPromotionbyId(id string) (promotion commands.Promotion_Command, err error)
	GetCustomerByTelephone_number(telephone_number string) (customer commands.Customer_command, err error)
	Delete_temp_order_by_id(orderID string) error
	Delete_temp_order_details_by_order_id(orderID string) error
}
