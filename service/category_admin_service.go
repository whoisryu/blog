package service

import (
	"blog/entity"
	"blog/model"
)

type CategoryAdminService interface {
	CreateCategory(req model.CreateCategoryRequest) entity.Category
}
