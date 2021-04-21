package service

import "blog/model"

type CategoryService interface {
	FindAll() (responses []model.CategoryResponse)
	FindByID(ID string) (response model.CategoryResponse)
}
