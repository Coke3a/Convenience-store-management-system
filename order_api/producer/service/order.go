package service_producer

import (
	"github.com/Coke3a/Convenience-store-management-system/order/producer/commands"
	repository_producer "github.com/Coke3a/Convenience-store-management-system/order/producer/repository"
)

type orderService struct {
	orderRepo repository_producer.OrderRepository	
}

func NewOrderService(orderRepo repository_producer.OrderRepository) OrderService {
	return orderService{orderRepo: orderRepo}
}

type OrderService interface {
	// CreateNewOrder() error
	// GetLastOrder() (order commands.Order_Command, err error)
	CreateNewOrderAndGetLastOrder() (order commands.Order_Command, err error)
	AddProduct(product_barcode string, order_id string)  (product commands.Product_Command, err error)
	DeleteOrder_Detail(id string) error
	Total_Price_Without_Promotion(order_id string) (total_price int, err error)
	FindCustomerByTelephone_number(order_id string, telephone_number string) (customer commands.Customer_command, err error)
	DeleteCustomer(order_id string) error
	FindPromotionwithConditions(order_id string) (promotions []commands.Promotion_Command, err error)
	FindPromotionByIdAndCalculate(order_id string, promotion_id string) (TotalWithPromo int, discount int, err error)	
	CalculatePromotionDiscountPrice(totalNopromo int, promotion commands.Promotion_Command) (TotalWithPromo int, discount int)
	CalculatePromotionDiscountPercent(totalNopromo int, promotion commands.Promotion_Command) (TotalWithPromo int, discount int)
	FindAllOrderDetailByOrderID(order_id string) (order_details []commands.Order_detail_Command , err error)
}
