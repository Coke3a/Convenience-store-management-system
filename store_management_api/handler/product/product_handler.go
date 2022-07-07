package handler_product

import "github.com/gofiber/fiber/v2"

func (h productHandler) CreateNewProduct(c *fiber.Ctx) error {
	request := ProductRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}
	err = h.productSev.AddNewProduct(request.Product_barcode, request.Name, request.Price)
	if err != nil {
		return err
	}
	response := fiber.Map{
		"product": request,
		"status":  "Created",
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h productHandler) FindAllProducts(c *fiber.Ctx) error {
	products, err := h.productSev.FindAllProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"products": products,
	}
	return c.JSON(response)
}

func (h productHandler) FindProductByid(c *fiber.Ctx) error {
	id := c.Params("id")


	product, err := h.productSev.FindProductById(id)
	if err != nil {
		return err
	}
	response := fiber.Map{
		"product": product,
	}
	return c.JSON(response)
}

func (h productHandler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")


	request := ProductRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	err = h.productSev.UpdateProduct(id, request.Product_barcode, request.Name, request.Price)
	if err != nil {
		return err
	}

	response := fiber.Map{
		"product": request,
		"status":  "Updated",
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h productHandler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")


	err := h.productSev.DeleteProduct(id)
	if err != nil {
		return err
	}
	response := fiber.Map{
		"product": id,
		"status":  "Deleted",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
