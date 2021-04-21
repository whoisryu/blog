package service

import "blog/model"

type PostService interface {
	ListPost(req model.ListPostRequest) (response []model.ListPostResponse)
	PostBySlug(slug string) (response model.ListPostResponse)
	CreatePost(req model.CreatePostRequest)
}
