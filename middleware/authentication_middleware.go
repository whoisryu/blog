package middleware

import (
	"blog/helper"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func TokenAuth() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		err := helper.TokenValid(c)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(helper.ResponseUnauthorized())

		}

		return c.Next()
	}
}
