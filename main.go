package main

import (
	"blog/config"
	"blog/controller"
	"blog/exception"
	"blog/repository"
	"blog/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	exception.PanicIfNeeded(err)

	//DATABASE
	db := config.MysqlConnection()

	//REPOSITORY
	userRepo := repository.NewUserRepo(db)
	postRepo := repository.NewPostRepo(db)
	categoryRepo := repository.NewCategoryRepo(db)
	commentRepo := repository.NewCommentRepo(db)

	//SERVICE
	userService := service.NewUserService(&userRepo)
	postService := service.NewPostService(&postRepo)
	categoryService := service.NewCategoryService(&categoryRepo)
	commentService := service.NewCommentService(&commentRepo)

	//CONTROLLER
	userController := controller.NewUserController(&userService)
	postController := controller.NewPostController(&postService)
	categoryController := controller.NewCategoryController(&categoryService)
	commentController := controller.NewCommentController(&commentService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	v1 := app.Group("/api/v1/blog")

	//ROUTE
	userController.Route(v1)
	postController.Route(v1)
	categoryController.Route(v1)
	commentController.Route(v1)

	err = app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
