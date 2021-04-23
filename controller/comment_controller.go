package controller

import (
	"blog/exception"
	"blog/helper"
	"blog/middleware"
	"blog/model"
	"blog/service"

	"github.com/gofiber/fiber/v2"
)

type CommentController struct {
	service service.CommentService
}

func NewCommentController(commentService *service.CommentService) CommentController {
	return CommentController{service: *commentService}
}

func (controller CommentController) Route(app fiber.Router) {
	commentRouter := app.Group("/comment")

	commentRouter.Get("/:id", middleware.TokenAuth(), controller.ListComment)
	commentRouter.Post("/:id", middleware.TokenAuth(), controller.CreateComment)
}

func (controller CommentController) ListComment(c *fiber.Ctx) error {
	postId := c.Params("id")

	payload := model.ListCommentRequest{
		PostID: postId,
	}

	response := controller.service.ListComment(payload)

	return c.Status(200).JSON(helper.ResponseSuccess(response))
}

func (controller CommentController) CreateComment(c *fiber.Ctx) error {
	au, err := helper.ExtractTokenMetadata(c)
	exception.PanicIfNeeded(err)
	postID := c.Params("id")
	userID := au.UserId

	comment := new(model.CommentRequest)
	comment.PostID = postID
	comment.CommentBy = userID

	if err := c.BodyParser(&comment); err != nil {
		exception.PanicIfNeeded(err)
	}

	response := controller.service.CreateComment(*comment)

	return c.Status(200).JSON(helper.ResponseSuccess(response))

}
