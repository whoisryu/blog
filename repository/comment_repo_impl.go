package repository

import (
	"blog/entity"
	"blog/exception"
	"blog/model"

	"gorm.io/gorm"
)

type commentRepoImpl struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepo {
	return &commentRepoImpl{db: db}
}

func (repo commentRepoImpl) CreateComment(comment entity.Comment) entity.Comment {
	err := repo.db.Create(&comment).Error

	exception.PanicIfNeeded(err)

	return comment
}

func (repo commentRepoImpl) ListComment(req model.ListCommentRequest) (response []entity.Comment) {

	err := repo.db.Where("post_id=?", req.PostID).Find(&response).Error
	exception.PanicIfNeeded(err)

	return response
}
