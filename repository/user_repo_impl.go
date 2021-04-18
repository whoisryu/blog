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

func (repo userRepoImpl) CreateUser(user entity.User) (entity.User, error) {
	err := repo.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo userRepoImpl) GetUserByPhone(phone string) (user entity.User, err error) {
	err = repo.db.Where("phone = ?", phone).Find(&user).Error
	exception.PanicIfNeeded(err)
	return user, err
}

func (repo userRepoImpl) GetUserByEmail(email string) (user entity.User, err error) {
	err = repo.db.Where("email = ?", email).Find(&user).Error

	exception.PanicIfNeeded(err)
	return user, nil
}
