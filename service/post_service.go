package service

import (
	"blog/entity"
	"blog/model"
)

type PostService interface {
	ListPost(req model.ListPostRequest) (response []model.ListPostResponse)
	ListPostByCategory(req model.ListPostByCategoryRequest) (response []model.ListPostResponse)
	PostBySlug(req model.PostBySlug) (response model.ListPostResponse)
	CreatePost(req model.CreatePostRequest) entity.Post
	UpdatePost(req model.UpdatePostRequest) entity.Post
	DeletePost(ID string)
}
