package handler_promotion

import "github.com/gofiber/fiber/v2"



func (h promotionHandler) CreateNewPromotionDiscountPrice(c *fiber.Ctx) error {
	request := PromotionRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err  
	}
	err = h.promotionServ.AddNewPromotionDiscountPrice(request.Name,request.PurchaseMin,request.Required_point,request.Discount)
	if err != nil {
		return err  
	}
	response := fiber.Map{
		"customer" : request,
		"status" : "Created",
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h promotionHandler) CreateNewPromotionDiscountPercent(c *fiber.Ctx) error {
	request := PromotionRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err  
	}
	err = h.promotionServ.AddNewPromotionDiscountPercent(request.Name,request.PurchaseMin,request.Required_point,request.Discount)
	if err != nil {
		return err  
	}
	response := fiber.Map{
		"customer" : request,
		"status" : "Created",
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}


func (h promotionHandler) FindAllPromotion(c *fiber.Ctx) error {
	promotions, err := h.promotionServ.FindAllPromotion()
	if err != nil {
		return err  
	}
	response := fiber.Map{
		"promotions" : promotions,
	}
	return c.JSON(response)
}

func (h promotionHandler) FindPromotionByid(c *fiber.Ctx) error {
	id := c.Params("id")

	promotion, err := h.promotionServ.FindPromotionById(id)
	if err != nil {
		return err  
	}
	response := fiber.Map{
		"promotion" : promotion,
	}
	return c.JSON(response)
}

func (h promotionHandler) UpdatePromotion(c *fiber.Ctx) error {
	id := c.Params("id")


	request := PromotionRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err  
	}

	err = h.promotionServ.UpdatePromotion(id,request.Name,request.PurchaseMin,request.Required_point,request.Discount)
	if err != nil {
		return err  
	}
	response := fiber.Map{
		"promotion" : request,
		"status" : "Updated",
	}

	return c.Status(fiber.StatusOK).JSON(response)
}



func (h promotionHandler) DeletePromotion(c *fiber.Ctx) error {
	id := c.Params("id")


	err := h.promotionServ.DeletePromotion(id)
	if err != nil {
		return err  
	}
	
	response := fiber.Map{
		"promotion" : id,
		"status" : "Deleted",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
