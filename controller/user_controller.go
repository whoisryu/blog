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
	userRoute.Post("/login", controller.Login)
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

func (controller UserController) Login(c *fiber.Ctx) error {
	user := new(model.LoginRequest)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.ResponseInternalError(err))
	}

	token, err := controller.UserService.Login(*user)

	if err != nil && err.Error() == "404" {
		return c.Status(http.StatusNotFound).JSON(helper.ResponseNotFound())
	}

	if err != nil && err.Error() == "401" {
		return c.Status(http.StatusUnauthorized).JSON(helper.ResponseUnauthorized())
	}

	return c.Status(http.StatusAccepted).JSON(helper.ResponseSuccess(token))
}
