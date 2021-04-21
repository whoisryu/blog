package repository

import (
	"blog/entity"
	"blog/exception"

	"gorm.io/gorm"
)

type userRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepoImpl{db: db}
}

func (repo userRepoImpl) CreateUser(user entity.User) entity.User {
	err := repo.db.Create(&user).Error

	exception.PanicIfNeeded(err)

	return user
}

func (repo userRepoImpl) GetUserByPhone(phone string) (user entity.User) {
	err := repo.db.Where("phone = ?", phone).Find(&user).Error
	exception.PanicIfNeeded(err)
	return user
}

func (repo userRepoImpl) GetUserByEmail(email string) (user entity.User) {
	err := repo.db.Where("email = ?", email).Find(&user).Error

	exception.PanicIfNeeded(err)
	return user
}

func (repo userRepoImpl) GetUserByID(user entity.User) entity.User {
	err := repo.db.First(&user).Error
	exception.PanicIfNeeded(err)

	return user
}

func (repo userRepoImpl) UpdateUser(user entity.User) entity.User {
	err := repo.db.Updates(&user).Error
	exception.PanicIfNeeded(err)

	return user
}
