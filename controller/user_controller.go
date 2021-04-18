package controller

import (
	"blog/helper"
	"blog/model"
	"blog/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{UserService: *userService}
}

func (controller *UserController) Route(app fiber.Router) {
	userRoute := app.Group("/user")

	userRoute.Post("/register", controller.RegisterUser)
}

func (controller *UserController) RegisterUser(c *fiber.Ctx) error {
	user := new(model.RegisterUserRequest)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.ResponseInternalError(err))
	}

	newUser, err := controller.UserService.RegisterUser(*user)

	if err != nil && err.Error() == "phone registered" {
		return c.Status(http.StatusBadRequest).JSON(helper.ResponseBadRequest(map[string]interface{}{"phone": "REGISTERED"}))
	}

	if err != nil && err.Error() == "email registered" {
		return c.Status(http.StatusBadRequest).JSON(helper.ResponseBadRequest(map[string]interface{}{"email": "REGISTERED"}))
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.ResponseInternalError(err))
	}

	return c.Status(http.StatusCreated).JSON(helper.ResponseSuccess(newUser))
}
