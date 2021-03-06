package controller

import (
	"blog/exception"
	"blog/helper"
	"blog/middleware"
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
	postRoute.Get("/mine", middleware.TokenAuth(), controller.ListPostUser)
	postRoute.Get("/:slug", controller.PostBySlug)
	postRoute.Get("/topic/:slug", controller.ListByCategory)
	postRoute.Post("/", middleware.TokenAuth(), controller.CreatePost)
	postRoute.Put("/:id", middleware.TokenAuth(), controller.UpdatePost)
	postRoute.Delete(":id", middleware.TokenAuth(), controller.DeletePost)
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

func (controller *PostController) ListPostUser(c *fiber.Ctx) error {
	au, _ := helper.ExtractTokenMetadata(c)

	getUserId := au.UserId

	userId := strconv.Itoa(int(getUserId))
	q := c.Query("q")
	getSort := c.Query("sort")

	sort, err := strconv.Atoi(getSort)

	if err != nil {
		sort = 0
	}

	payload := model.ListPostRequestMine{
		SortBy: uint(sort),
		Q:      q,
		UserID: userId,
	}

	response := controller.PostService.ListPostUser(payload)

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
	au, _ := helper.ExtractTokenMetadata(c)

	var userID = ""

	if au != nil {
		userID = strconv.Itoa(int(au.UserId))

	}

	request := model.PostBySlug{
		Slug:   slug,
		UserID: userID,
	}

	response := controller.PostService.PostBySlug(request)

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

	return c.Status(201).JSON(helper.ResponseSuccess(newPost))
}

func (controller *PostController) UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")

	post := new(model.UpdatePostRequest)

	post.ID = id

	if err := c.BodyParser(&post); err != nil {
		exception.PanicIfNeeded(err)
	}

	updatedPost := controller.PostService.UpdatePost(*post)

	return c.Status(201).JSON(helper.ResponseSuccess(updatedPost))
}

func (controller *PostController) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")

	controller.PostService.DeletePost(id)

	return c.Status(201).JSON(helper.ResponseSuccess(struct{}{}))
}
