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

	//SERVICE
	userService := service.NewUserService(&userRepo)
	postService := service.NewPostService(&postRepo)

	//CONTROLLER
	userController := controller.NewUserController(&userService)
	postController := controller.NewPostController(&postService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	v1 := app.Group("/api/v1/blog")

	//ROUTE
	userController.Route(v1)
	postController.Route(v1)

	err = app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
