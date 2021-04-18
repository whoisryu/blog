package repository

import "blog/entity"

type UserRepo interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUserByPhone(phone string) (user entity.User, err error)
	GetUserByEmail(email string) (user entity.User, err error)
}
