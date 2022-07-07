package handler_producer

import (
	"github.com/Coke3a/Convenience-store-management-system/order/producer/commands"
	"github.com/gofiber/fiber/v2"
)

//get
func (h orderHandler) CreateNewOrderAndGetLastOrder(c *fiber.Ctx) error {
	order, err := h.orderSev.CreateNewOrderAndGetLastOrder()
	if err != nil {
		return err
	}
	response := fiber.Map{
		"order_id": order.ID,
		"status":   "created",
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h orderHandler) CreateOrder_DetailAndAddProduct(c *fiber.Ctx) error {

	orderID := c.Params("id")

	request := commands.Product_Command{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	product, err := h.orderSev.AddProduct(request.Product_barcode, orderID)
	if err != nil {
		return err
	}

	response := fiber.Map{
		"product": product,
		"status":  "created",
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}


func (h orderHandler) FindAllProductsByOrderID(c *fiber.Ctx) error {
	orderID := c.Params("id")

	order_details , err := h.orderSev.FindAllOrderDetailByOrderID(orderID)
	if err != nil {
		return err
	}
	return c.JSON(order_details)

}




func (h orderHandler) DeleteOrder_Detail(c *fiber.Ctx) error {
	request := commands.Product_Command{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}
	h.orderSev.DeleteOrder_Detail(request.ID)

	response := fiber.Map{
		"product": request.ID,
		"status":  "deleted",
	}

	return c.JSON(response)
}

//get
func (h orderHandler) Total_Price_Without_Promotion(c *fiber.Ctx) error {
	orderID := c.Params("id")
	total, err := h.orderSev.Total_Price_Without_Promotion(orderID)
	if err != nil {
		return err
	}
	response := fiber.Map{
		"total": total,
	}
	return c.JSON(response)
}

//post
func (h orderHandler) FindCustomerByTelephone_number(c *fiber.Ctx) error {
	orderID := c.Params("id")
	request := commands.Customer_command{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}
	customer, err := h.orderSev.FindCustomerByTelephone_number(orderID, request.Telephone_number)
	if err != nil {
		return c.JSON(fiber.Map{"status":"customer not found"})
	}
	response := fiber.Map{
		"customer": customer,
	}
	return c.JSON(response)
}

//put
func (h orderHandler) DeleteCustomer(c *fiber.Ctx) error {
	orderID := c.Params("id")
	err := h.orderSev.DeleteCustomer(orderID)
	if err != nil {
		return err
	}
	response := fiber.Map{
		"customer": orderID,
		"status":   "deleted",
	}
	return c.JSON(response)
}

//get
func (h orderHandler) FindPromotionwithConditions(c *fiber.Ctx) error {
	orderID := c.Params("id")
	promotions, err := h.orderSev.FindPromotionwithConditions(orderID)
	if err != nil {
		return err
	}
	return c.JSON(promotions)
}

//post
func (h orderHandler) FindPromotionByIdAndCalculate(c *fiber.Ctx) error {
	orderID := c.Params("id")
	promotionID := c.Params("promotion_id")
	// request := commands.Promotion_Command{}
	// err := c.BodyParser(&request)
	// if err != nil {
	// 	return err
	// }
	total, discount, err := h.orderSev.FindPromotionByIdAndCalculate(orderID, promotionID)
	if err != nil {
		return err
	}
	response := fiber.Map{
		"total":    total,
		"discount": discount,
	}
	return c.JSON(response)

}

//put
func (h orderHandler) SubmitOrderAndOrderDetail(c *fiber.Ctx) error {
	orderID := c.Params("id")
	err := h.produceService.SubmitOrder(orderID)
	if err != nil {
		return err
	}
	err = h.produceService.LoopForSubmitOrder_Detail(orderID)
	if err != nil {
		return err
	}
	response := fiber.Map{
		"order_id":    orderID,
		"status": "process successfully completed",
	}

	err = h.produceService.DeleteOrderIdAndOrder_detailByOrderID(orderID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
