package controller

import (
	"blog/exception"
	"blog/model"
	"blog/service"

	"github.com/gofiber/fiber/v2"
)

type CategoryAdminController struct {
	CategoryAdminService service.CategoryAdminService
}

func NewCategoryAdminController(service *service.CategoryAdminService) CategoryAdminController {
	return CategoryAdminController{CategoryAdminService: *service}
}

func (controller CategoryAdminController) Route(app fiber.Router) {
	router := app.Group("/category")

	router.Post("/", controller.CreateCategory)
}

func (controller CategoryAdminController) CreateCategory(c *fiber.Ctx) error {
	category := new(model.CreateCategoryRequest)

	if err := c.BodyParser(&category); err != nil {
		exception.PanicIfNeeded(err)
	}

	newCategory := controller.CategoryAdminService.CreateCategory(*category)

	return c.Status(200).JSON(newCategory)
}
