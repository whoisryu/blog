package service

import (
	"blog/entity"
	"blog/exception"
	"blog/model"
	"blog/repository"
	"blog/validation"
	"strconv"
	"strings"
	"time"
)

type postServiceImpl struct {
	postRepository repository.PostRepo
}

func NewPostService(postRepo *repository.PostRepo) PostService {
	return &postServiceImpl{postRepository: *postRepo}
}

func (service postServiceImpl) ListPost(req model.ListPostRequest) (response []model.ListPostResponse) {
	response = service.postRepository.ListPost(req)

	return response
}

func (service postServiceImpl) PostBySlug(slug string) (response model.ListPostResponse) {
	response = service.postRepository.PostBySlug(slug)

	return response
}

func (service postServiceImpl) CreatePost(req model.CreatePostRequest) entity.Post {
	validation.ValidatePost(req)
	slug := strings.ReplaceAll(strings.ToLower(req.Title), " ", "-")

	categories := []entity.Category{}

	for _, category := range req.Categories {
		categories = append(categories, entity.Category{ID: category})
	}

	post := entity.Post{
		AuthorId:    req.AuthorId,
		Title:       req.Title,
		Slug:        slug,
		Content:     req.Content,
		IsPublished: req.IsPublished,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Categories:  categories,
	}

	if req.IsPublished {
		post.PublishedAt = time.Now()
	}

	newPost := service.postRepository.CreatePost(post)
	return newPost
}

func (service postServiceImpl) UpdatePost(req model.UpdatePostRequest) entity.Post {
	validation.ValidateUpdatePost(req)
	id, _ := strconv.Atoi(req.ID)
	slug := strings.ReplaceAll(strings.ToLower(req.Title), " ", "-")

	categories := []entity.Category{}

	for _, category := range req.Categories {
		categories = append(categories, entity.Category{ID: category})
	}

	service.postRepository.PostByID(uint(id))

	post := entity.Post{
		ID:          uint(id),
		AuthorId:    req.AuthorId,
		Title:       req.Title,
		Slug:        slug,
		Content:     req.Content,
		IsPublished: req.IsPublished,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Categories:  categories,
	}

	if req.IsPublished {
		post.PublishedAt = time.Now()
	}

	updatedPost := service.postRepository.UpdatePost(post)

	return updatedPost
}

func (service postServiceImpl) DeletePost(ID string) {
	id, err := strconv.Atoi(ID)
	exception.PanicIfNeeded(err)
	service.postRepository.PostByID(uint(id))

	service.postRepository.DeletePost(uint(id))
}

func (service postServiceImpl) ListPostByCategory(req model.ListPostByCategoryRequest) (response []model.ListPostResponse) {
	response = service.postRepository.ListByCategory(req)

	return response
}
