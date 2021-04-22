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
	categoryAdminService := service.NewCategoryAdminService(&categoryRepo)

	//CONTROLLER
	userController := controller.NewUserController(&userService)
	postController := controller.NewPostController(&postService)
	categoryController := controller.NewCategoryController(&categoryService)
	commentController := controller.NewCommentController(&commentService)
	categoryAdminController := controller.NewCategoryAdminController(&categoryAdminService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	v1 := app.Group("/api/v1/blog")

	v1Admin := app.Group("/api/v1/admin/blog")

	// v1Admin.Use(basicauth.New(basicauth.Config{
	// 	Users: map[string]string{
	// 		"admin": "admin",
	// 	},
	// }))

	//ROUTE
	userController.Route(v1)
	postController.Route(v1)
	categoryController.Route(v1)
	commentController.Route(v1)

	//ROUTE_ADMIN
	categoryAdminController.Route(v1Admin)

	err = app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
