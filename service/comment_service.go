package service

import (
	"blog/entity"
	"blog/model"
)

type CommentService interface {
	CreateComment(req model.CommentRequest) (comment entity.Comment)
	ListComment(req model.ListCommentRequest) (response []model.ListCommentResponse)
}
