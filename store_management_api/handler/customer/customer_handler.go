package handler_customer

import "github.com/gofiber/fiber/v2"

func (h customerHandler) CreateNewCustomer(c *fiber.Ctx) error {
	request := CustomerRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}
	err = h.customerSev.AddNewCustomer(request.First_name, request.Last_name, request.Telephone_number)
	if err != nil {
		return err
	}

	response := fiber.Map{
		"customer": request,
		"status":   "Created",
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h customerHandler) FindAllCustomer(c *fiber.Ctx) error {
	customers, err := h.customerSev.FindAllCustomer()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"customers": customers,
	}
	return c.JSON(response)
}

func (h customerHandler) FindCustomerByid(c *fiber.Ctx) error {
	id := c.Params("id")


	customer, err := h.customerSev.FindCustomerById(id)
	if err != nil {
		return err
	}

	response := fiber.Map{
		"customer": customer,
	}
	return c.JSON(response)
}

func (h customerHandler) UpdateCustomer(c *fiber.Ctx) error {
	id := c.Params("id")


	request := CustomerRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	err = h.customerSev.UpdateCustomer(id, request.First_name, request.Last_name, request.Telephone_number)
	if err != nil {
		return err
	}
	response := fiber.Map{
		"customer": request,
		"status":   "Updated",
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h customerHandler) DeleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.customerSev.DeleteCustomer(id)
	if err != nil {
		return err
	}
	response := fiber.Map{
		"customer": id,
		"status":   "Deleted",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
