package service

import (
	"blog/entity"
	"blog/exception"
	"blog/helper"
	"blog/model"
	"blog/repository"
	"blog/validation"
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
	validation.ValidateRegisterUser(req)

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

	newUser := service.UserRepo.CreateUser(user)

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
	validation.ValidateLoginUser(req)
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

func (service userServiceImpl) UpdateProfile(req model.UpdateProfileRequest) (entity.User, error) {
	validation.ValidateUpdate(req)
	userID, err := strconv.Atoi(req.ID)
	exception.PanicIfNeeded(err)

	service.UserRepo.GetUserByID(entity.User{ID: uint(userID)})

	checkUserPhone := service.UserRepo.GetUserByPhone(req.Phone)
	if checkUserPhone.ID != 0 {
		return entity.User{}, errors.New("phone registered")
	}

	checkUserEmail := service.UserRepo.GetUserByEmail(req.Email)
	if checkUserEmail.ID != 0 {
		return entity.User{}, errors.New("email registered")
	}

	id, err := strconv.Atoi(req.ID)
	exception.PanicIfNeeded(err)
	user := entity.User{
		ID:      uint(id),
		Phone:   req.Phone,
		Email:   req.Email,
		Profile: req.Profile,
	}

	newUser := service.UserRepo.UpdateUser(user)

	return newUser, nil

}
