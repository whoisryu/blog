package service

import (
	"blog/entity"
	"blog/exception"
	"blog/model"
	"blog/repository"
	"strconv"
	"strings"
)

type categoryAdminServiceImpl struct {
	repo repository.CategoryRepo
}

func NewCategoryAdminService(repository *repository.CategoryRepo) CategoryAdminService {
	return &categoryAdminServiceImpl{repo: *repository}
}

func (service categoryAdminServiceImpl) CreateCategory(req model.CreateCategoryRequest) entity.Category {
	slug := strings.ReplaceAll(strings.ToLower(req.Title), " ", "-")

	parentID, err := strconv.Atoi(req.ParentID)
	exception.PanicIfNeeded(err)

	category := entity.Category{
		ParentID: uint(parentID),
		Title:    req.Title,
		Slug:     slug,
	}

	newCategory := service.repo.CreateCategory(category)

	return newCategory
}
