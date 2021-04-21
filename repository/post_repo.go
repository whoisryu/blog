package repository

import (
	"blog/entity"
	"blog/model"
)

type PostRepo interface {
	ListPost(req model.ListPostRequest) (response []model.ListPostResponse)
	PostBySlug(slug string) (response model.ListPostResponse)
	PostByID(ID uint) (post entity.Post)
	CreatePost(post entity.Post) entity.Post
	UpdatePost(post entity.Post) entity.Post
	DeletePost(ID uint)
}
