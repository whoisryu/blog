package controller

import (
	"blog/exception"
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

	postRoute.Get("/", controller.ListPost)
	postRoute.Get("/:slug", controller.PostBySlug)
	postRoute.Get("/topic/:slug", controller.ListByCategory)
	postRoute.Post("/", controller.CreatePost)
	postRoute.Put("/:id", controller.UpdatePost)
	postRoute.Delete(":id", controller.DeletePost)
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

	return c.Status(http.StatusOK).JSON(helper.ResponseSuccess(response))
}

func (controller *PostController) ListByCategory(c *fiber.Ctx) error {
	slug := c.Params("slug")

	request := model.ListPostByCategoryRequest{
		Slug: slug,
	}

	response := controller.PostService.ListPostByCategory(request)

	return c.Status(http.StatusOK).JSON(helper.ResponseSuccess(response))
}

func (controller *PostController) PostBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")

	response := controller.PostService.PostBySlug(slug)

	if response.ID == 0 {
		return c.Status(http.StatusNotFound).JSON(helper.ResponseNotFound())
	}

	return c.Status(http.StatusOK).JSON(helper.ResponseSuccess(response))

}

func (controller *PostController) CreatePost(c *fiber.Ctx) error {
	au, err := helper.ExtractTokenMetadata(c)
	exception.PanicIfNeeded(err)

	post := new(model.CreatePostRequest)

	post.AuthorId = uint(au.UserId)

	if err := c.BodyParser(&post); err != nil {
		exception.PanicIfNeeded(err)
	}

	newPost := controller.PostService.CreatePost(*post)

	return c.Status(http.StatusOK).JSON(helper.ResponseSuccess(newPost))
}

func (controller *PostController) UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")

	post := new(model.UpdatePostRequest)

	post.ID = id

	if err := c.BodyParser(&post); err != nil {
		exception.PanicIfNeeded(err)
	}

	updatedPost := controller.PostService.UpdatePost(*post)

	return c.Status(http.StatusOK).JSON(helper.ResponseSuccess(updatedPost))
}

func (controller *PostController) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")

	controller.PostService.DeletePost(id)

	return c.Status(http.StatusOK).JSON(helper.ResponseSuccess(struct{}{}))
}
