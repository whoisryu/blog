package repository

import (
	"blog/entity"
)

type CategoryRepo interface {
	FindAll() (categories []entity.Category)
	FindByID(ID uint) (category entity.Category)
	CreateCategory(category entity.Category) entity.Category
}
