package exception

import (
	"blog/model"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, ok := err.(ValidationError)
	if ok {
		var obj interface{}
		_ = json.Unmarshal([]byte(err.Error()), &obj)
		return ctx.Status(400).JSON(model.Response{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   struct{}{},
			Error:  obj,
		})
	}

	return ctx.Status(500).JSON(model.Response{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}
