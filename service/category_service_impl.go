package service

import (
	"blog/exception"
	"blog/model"
	"blog/repository"
	"strconv"
)

type categoryServiceImpl struct {
	categoryRepo repository.CategoryRepo
}

func NewCategoryService(repo *repository.CategoryRepo) CategoryService {
	return &categoryServiceImpl{categoryRepo: *repo}
}

func buildCategory(ids []model.CategoryResponse, relations map[uint][]model.CategoryResponse) []model.CategoryResponse {
	categories := make([]model.CategoryResponse, len(ids))
	for i, category := range ids {
		c := model.CategoryResponse{
			ID:       category.ID,
			ParentID: category.ParentID,
			Slug:     category.Slug,
			Title:    category.Title,
		}
		if childs, ok := relations[category.ID]; ok {
			c.Child = buildCategory(childs, relations)
		}
		categories[i] = c
	}

	return categories
}

func (service categoryServiceImpl) FindAll() (responses []model.CategoryResponse) {
	datas := service.categoryRepo.FindAll()
	dataset := []model.CategoryResponse{}
	for _, data := range datas {
		dataset = append(dataset, model.CategoryResponse{
			ID:       data.ID,
			ParentID: data.ParentID,
			Title:    data.Title,
			Slug:     data.Slug,
		})
	}

	relations := make(map[uint][]model.CategoryResponse)

	for _, relation := range dataset {
		child, parent := relation, relation.ParentID
		relations[parent] = append(relations[parent], child)
	}

	categories := buildCategory(relations[0], relations)

	return categories
}

func (service categoryServiceImpl) FindByID(ID string) (response model.CategoryResponse) {
	id, err := strconv.Atoi(ID)
	exception.PanicIfNeeded(err)

	data := service.categoryRepo.FindByID(uint(id))

	response = model.CategoryResponse{
		ID:       data.ID,
		ParentID: data.ParentID,
		Title:    data.Title,
		Slug:     data.Slug,
	}

	return response
}
