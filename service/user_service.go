package service

import (
	"blog/entity"
	"blog/model"
)

type UserService interface {
	RegisterUser(req model.RegisterUserRequest) (*model.TokenResponse, error)
	Login(req model.LoginRequest) (*model.TokenResponse, error)
	UpdateProfile(req model.UpdateProfileRequest) (entity.User, error)
}
