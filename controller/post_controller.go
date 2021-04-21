package controller

import (
	"blog/helper"
	"blog/model"
	"blog/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PostController struct {
	PostService service.PostService
}

func NewPostController(postService *service.PostService) PostController {
	return PostController{PostService: *postService}
}

func (controller *PostController) Route(app fiber.Router) {
	postRoute := app.Group("/post")

	postRoute.Get("/list", controller.ListPost)
	postRoute.Get("/list/:slug", controller.PostBySlug)
}

func (controller *PostController) ListPost(c *fiber.Ctx) error {
	q := c.Query("q")
	getSort := c.Query("sort")

	sort, err := strconv.Atoi(getSort)

	if err != nil {
		sort = 0
	}

	payload := model.ListPostRequest{
		SortBy: uint(sort),
		Q:      q,
	}

	response := controller.PostService.ListPost(payload)

	return c.Status(http.StatusAccepted).JSON(helper.ResponseSuccess(response))
}

func (controller *PostController) PostBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	response := controller.PostService.PostBySlug(slug)

	if response.ID == 0 {
		return c.Status(http.StatusNotFound).JSON(helper.ResponseNotFound())
	}

	return c.Status(http.StatusAccepted).JSON(helper.ResponseSuccess(response))

}
