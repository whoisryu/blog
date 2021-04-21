package repository

import (
	"blog/entity"
	"blog/exception"

	"gorm.io/gorm"
)

type categoryRepoImpl struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) CategoryRepo {
	return &categoryRepoImpl{db: db}
}

func (repo categoryRepoImpl) FindAll() (categories []entity.Category) {
	err := repo.db.Find(&categories).Error
	exception.PanicIfNeeded(err)

	return categories
}

func (repo categoryRepoImpl) FindByID(ID uint) (category entity.Category) {
	err := repo.db.First(&category).Error
	exception.PanicIfNeeded(err)

	return category
}
