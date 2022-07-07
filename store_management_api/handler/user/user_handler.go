package handler_user

import "github.com/gofiber/fiber/v2"


func (h userHandler) CreateNewUser(c *fiber.Ctx) error {
	request := UserRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}
	err = h.userServ.AddNewUser(request.Username,request.Password,request.Email, request.First_name, request.Last_name)
	if err != nil {
		return err
	}
	response := fiber.Map{
		"customer": request,
		"status":   "Created",
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h userHandler) FindAllUser(c *fiber.Ctx) error {
	customers, err := h.userServ.FindAllUsers()
	if err != nil {
		return err
	}
	response := fiber.Map{
		"customers": customers,
	}
	return c.JSON(response)
}

func (h userHandler) FindUserByid(c *fiber.Ctx) error {
	id := c.Params("id")
	customer, err := h.userServ.FindUserByID(id)
	if err != nil {
		return err
	}
	response := fiber.Map{
		"customer": customer,
	}
	return c.JSON(response)
}

func (h userHandler) UpdateUser(c *fiber.Ctx) error {

	id := c.Params("id")


	request := UserRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	err = h.userServ.UpdateUser(id, request.Username,request.Password,request.Email, request.First_name, request.Last_name)
	if err != nil {
		return err
	}
	response := fiber.Map{
		"customer": request,
		"status":   "Updated",
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h userHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.userServ.DeleteUser(id)
	if err != nil {
		return err
	}
	response := fiber.Map{
		"customer": id,
		"status":   "Deleted",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
