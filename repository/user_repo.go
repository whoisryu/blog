package repository

import "blog/entity"

type UserRepo interface {
	CreateUser(user entity.User) entity.User
	GetUserByPhone(phone string) (user entity.User)
	GetUserByEmail(email string) (user entity.User)
	GetUserByID(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
}
