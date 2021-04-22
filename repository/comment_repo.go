package repository

import (
	"blog/entity"
	"blog/model"
)

type CommentRepo interface {
	CreateComment(comment entity.Comment) entity.Comment
	ListComment(req model.ListCommentRequest) (response []entity.Comment)
}
