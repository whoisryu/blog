package helper

import (
	"blog/model"
	"fmt"
)

func ResponseSuccess(data interface{}) model.Response {
	response := model.Response{
		Code:   201,
		Status: "OK",
		Data:   data,
		Error:  struct{}{},
	}

	return response
}

func ResponseBadRequest(errors interface{}) model.Response {
	response := model.Response{
		Code:   400,
		Status: "Bad Request",
		Data:   struct{}{},
		Error:  errors,
	}

	return response
}

func ResponseInternalError(errors interface{}) model.Response {
	fmt.Print(errors)

	response := model.Response{
		Code:   500,
		Status: "Server Errors",
		Data:   struct{}{},
		Error:  model.GeneralError{General: "INTERNAL_ERROR"},
	}

	return response
}

func ResponseNotFound() model.Response {
	response := model.Response{
		Code:   404,
		Status: "Not Found",
		Data:   struct{}{},
		Error:  model.GeneralError{General: "NOT_FOUND"},
	}

	return response
}

func ResponseUnauthorized() model.Response {
	response := model.Response{
		Code:   401,
		Status: "Unauthorized",
		Data:   struct{}{},
		Error:  model.GeneralError{General: "UNAUTORIZED"},
	}

	return response
}
