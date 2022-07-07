package handler_producer

import (
	service_producer "github.com/Coke3a/Convenience-store-management-system/order/producer/service"
	"github.com/gofiber/fiber/v2"
)

type orderHandler struct {
	orderSev service_producer.OrderService
	produceService service_producer.ProduceService
}

func NewOrderHandler(orderSev service_producer.OrderService, produceService service_producer.ProduceService) OrderHandler {
	return orderHandler{orderSev, produceService}
}

type OrderHandler interface {
	CreateNewOrderAndGetLastOrder(c *fiber.Ctx) error
	CreateOrder_DetailAndAddProduct(c *fiber.Ctx) error
	DeleteOrder_Detail(c *fiber.Ctx) error
	Total_Price_Without_Promotion(c *fiber.Ctx) error
	FindCustomerByTelephone_number(c *fiber.Ctx) error
	DeleteCustomer(c *fiber.Ctx) error
	FindPromotionwithConditions(c *fiber.Ctx) error
	FindPromotionByIdAndCalculate(c *fiber.Ctx) error
	SubmitOrderAndOrderDetail(c *fiber.Ctx) error	
	FindAllProductsByOrderID(c *fiber.Ctx) error
}
