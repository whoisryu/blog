package service

import (
	"blog/model"
)

type UserService interface {
	RegisterUser(req model.RegisterUserRequest) (*model.TokenResponse, error)
}
