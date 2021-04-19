package service

import (
	"blog/entity"
	"blog/helper"
	"blog/model"
	"blog/repository"
	"errors"
	"strconv"
	"time"
)

type userServiceImpl struct {
	UserRepo repository.UserRepo
}

func NewUserService(userRepo *repository.UserRepo) UserService {
	return &userServiceImpl{UserRepo: *userRepo}
}

func (service userServiceImpl) RegisterUser(req model.RegisterUserRequest) (*model.TokenResponse, error) {

	checkUser := service.UserRepo.GetUserByPhone(req.Phone)
	if checkUser.ID != 0 {
		return nil, errors.New("phone registered")
	}

	checkUser = service.UserRepo.GetUserByEmail(req.Email)
	if checkUser.ID != 0 {
		return nil, errors.New("email registered")
	}

	user := entity.User{
		UserName:     req.UserName,
		Phone:        req.Phone,
		Email:        req.Email,
		PasswordHash: helper.HashPwd([]byte(req.Password)),
		RegisteredAt: time.Now(),
		LastLogin:    time.Now(),
		Profile:      req.Profile,
	}

	newUser, err := service.UserRepo.CreateUser(user)

	if err != nil {
		return nil, err
	}
	payloadToken := model.JwtPayload{
		UserID:   strconv.Itoa(int(newUser.ID)),
		UserName: newUser.UserName,
	}

	ts, err := helper.CreateToken(payloadToken)
	if err != nil {
		return nil, err
	}

	saveErr := helper.CreateAuth(int64(newUser.ID), ts)
	if saveErr != nil {
		return nil, saveErr
	}

	tokens := model.TokenResponse{
		AccessToken:  ts.AccessToken,
		RefreshToken: ts.RefreshToken,
	}
	return &tokens, nil
}

func (service userServiceImpl) Login(req model.LoginRequest) (*model.TokenResponse, error) {
	user := service.UserRepo.GetUserByEmail(req.Email)

	if user.ID == 0 {
		return nil, errors.New("404")
	}

	isValid := helper.ComparePasswords(user.PasswordHash, []byte(req.Password))

	if isValid {
		payloadToken := model.JwtPayload{
			UserID:   strconv.Itoa(int(user.ID)),
			UserName: user.UserName,
		}
		token, err := helper.CreateToken(payloadToken)
		if err != nil {
			return nil, err
		}

		saveErr := helper.CreateAuth(int64(user.ID), token)
		if saveErr != nil {
			return nil, saveErr
		}

		tokenResp := model.TokenResponse{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		}

		return &tokenResp, nil
	}

	return nil, errors.New("401")

}
