package service

import (
	"blog/entity"
	"blog/model"
	"blog/repository"
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

func (service postServiceImpl) CreatePost(req model.CreatePostRequest) {
	slug := strings.ReplaceAll(strings.ToLower(req.Title), " ", "-")

	post := entity.Post{
		AuthorId:    req.AuthorId,
		Title:       req.Title,
		Slug:        slug,
		Content:     req.Content,
		IsPublished: req.IsPublished,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if req.IsPublished {
		post.PublishedAt = time.Now()
	}

	service.postRepository.CreatePost(post)
}
