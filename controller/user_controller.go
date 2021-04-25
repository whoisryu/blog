package controller

import (
	"blog/exception"
	"blog/helper"
	"blog/middleware"
	"blog/model"
	"blog/service"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{UserService: *userService}
}

func (controller *UserController) Route(app fiber.Router) {
	app.Post("/token/refresh", controller.RefreshToken)
	userRoute := app.Group("/user")
	userRoute.Post("/register", controller.RegisterUser)
	userRoute.Post("/login", controller.Login)
	userRoute.Post("/logout", middleware.TokenAuth(), controller.Logout)
	userRoute.Put("/", middleware.TokenAuth(), controller.UpdateProfile)
}

func (controller *UserController) RegisterUser(c *fiber.Ctx) error {
	user := new(model.RegisterUserRequest)

	if err := c.BodyParser(&user); err != nil {
		exception.PanicIfNeeded(err)
	}

	newUser, err := controller.UserService.RegisterUser(*user)

	if err != nil && err.Error() == "phone registered" {
		return c.Status(http.StatusBadRequest).JSON(helper.ResponseBadRequest(map[string]interface{}{"phone": "REGISTERED"}))
	}

	if err != nil && err.Error() == "email registered" {
		return c.Status(http.StatusBadRequest).JSON(helper.ResponseBadRequest(map[string]interface{}{"email": "REGISTERED"}))
	}

	return c.Status(http.StatusCreated).JSON(helper.ResponseSuccess(newUser))
}

func (controller UserController) UpdateProfile(c *fiber.Ctx) error {
	au, err := helper.ExtractTokenMetadata(c)
	exception.PanicIfNeeded(err)

	userID := strconv.Itoa(int(au.UserId))

	user := new(model.UpdateProfileRequest)
	user.ID = userID
	if err := c.BodyParser(&user); err != nil {
		exception.PanicIfNeeded(err)
	}

	newUser, err := controller.UserService.UpdateProfile(*user)

	if err != nil && err.Error() == "phone registered" {
		return c.Status(http.StatusBadRequest).JSON(helper.ResponseBadRequest(map[string]interface{}{"phone": "REGISTERED"}))
	}

	if err != nil && err.Error() == "email registered" {
		return c.Status(http.StatusBadRequest).JSON(helper.ResponseBadRequest(map[string]interface{}{"email": "REGISTERED"}))
	}

	return c.Status(201).JSON(helper.ResponseSuccess(newUser))
}

func (controller UserController) Login(c *fiber.Ctx) error {
	user := new(model.LoginRequest)

	if err := c.BodyParser(&user); err != nil {
		exception.PanicIfNeeded(err)
	}
	token, err := controller.UserService.Login(*user)
	if err != nil && err.Error() == "404" {
		return c.Status(http.StatusNotFound).JSON(helper.ResponseNotFound())
	}

	if err != nil && err.Error() == "401" {
		return c.Status(http.StatusUnauthorized).JSON(helper.ResponseUnauthorized())
	}

	return c.Status(http.StatusOK).JSON(helper.ResponseSuccess(token))
}

func (controller UserController) Logout(c *fiber.Ctx) error {
	au, err := helper.ExtractTokenMetadata(c)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusUnauthorized).JSON(helper.ResponseUnauthorized())
	}

	deleted, delErr := helper.DeleteAuth(au.AccessUuid)
	if delErr != nil && deleted == 0 {
		return c.Status(http.StatusUnauthorized).JSON(helper.ResponseUnauthorized())
	}

	return c.Status(http.StatusOK).JSON(helper.ResponseSuccess(struct{}{}))
}

func (controller UserController) RefreshToken(c *fiber.Ctx) error {
	req := new(model.RefreshTokenRequest)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.ResponseInternalError(err))
	}

	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(helper.ResponseUnauthorized())
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(helper.ResponseUnauthorized())
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string)
		if !ok {
			return c.Status(http.StatusUnprocessableEntity).Send([]byte(err.Error()))
		}

		userId, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
		userName := fmt.Sprintf("%v", claims["username"])
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(helper.ResponseInternalError("ERROR OCCURED"))
		}

		deleted, delErr := helper.DeleteAuth(refreshUuid)
		if delErr != nil && deleted == 0 {
			return c.Status(http.StatusUnauthorized).JSON(helper.ResponseUnauthorized())
		}

		payloadToken := model.JwtPayload{
			UserID:   strconv.Itoa(int(userId)),
			UserName: userName,
		}

		ts, createErr := helper.CreateToken(payloadToken)
		if createErr != nil {
			return c.Status(http.StatusForbidden).JSON(createErr.Error())
		}

		saveErr := helper.CreateAuth(int64(userId), ts)
		if saveErr != nil {
			return c.Status(http.StatusForbidden).JSON(saveErr.Error())
		}

		response := model.TokenResponse{
			AccessToken:  ts.AccessToken,
			RefreshToken: ts.RefreshToken,
		}

		return c.Status(http.StatusOK).JSON(helper.ResponseSuccess(response))

	} else {
		return c.Status(http.StatusUnauthorized).JSON(helper.ResponseUnauthorized())
	}
}
