package controller

import (
	"blog/helper"
	"blog/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService *service.CategoryService) CategoryController {
	return CategoryController{CategoryService: *categoryService}
}

func (controller *CategoryController) Route(app fiber.Router) {
	categoryRoute := app.Group("/category")

	categoryRoute.Get("/", controller.FindAll)
	categoryRoute.Get("/:id", controller.FindByID)
}

func (controller *CategoryController) FindAll(c *fiber.Ctx) error {
	response := controller.CategoryService.FindAll()

	return c.Status(http.StatusOK).JSON(helper.ResponseSuccess(response))
}

func (controller *CategoryController) FindByID(c *fiber.Ctx) error {
	ID := c.Params("id")

	response := controller.CategoryService.FindByID(ID)

	return c.Status(http.StatusOK).JSON(helper.ResponseSuccess(response))
}
