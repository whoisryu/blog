package service

import (
	"blog/entity"
	"blog/exception"
	"blog/model"
	"blog/repository"
	"blog/validation"
	"strconv"
	"time"
)

type commentServiceImpl struct {
	repo repository.CommentRepo
}

func NewCommentService(repo *repository.CommentRepo) CommentService {
	return &commentServiceImpl{repo: *repo}
}

func (service commentServiceImpl) CreateComment(req model.CommentRequest) (comment entity.Comment) {
	validation.ValidateComment(req)
	postID, err := strconv.Atoi(req.PostID)
	exception.PanicIfNeeded(err)

	parentID, err := strconv.Atoi(req.ParentID)
	exception.PanicIfNeeded(err)

	comment = entity.Comment{
		PostID:      uint(postID),
		ParentID:    uint(parentID),
		CreatedAt:   time.Now(),
		Content:     req.Content,
		CommentByID: uint(req.CommentBy),
	}

	comment = service.repo.CreateComment(comment)

	return comment
}

func buildComment(ids []model.ListCommentResponse, relations map[uint][]model.ListCommentResponse) []model.ListCommentResponse {
	comments := make([]model.ListCommentResponse, len(ids))
	for i, comment := range ids {
		c := model.ListCommentResponse{
			PostID:    comment.PostID,
			ParentID:  comment.ParentID,
			Content:   comment.Content,
			CommentBy: comment.CommentBy,
		}

		if childs, ok := relations[comment.ID]; ok {
			c.Child = buildComment(childs, relations)
		}
		comments[i] = c
	}

	return comments
}

func (service commentServiceImpl) ListComment(req model.ListCommentRequest) (comment []model.ListCommentResponse) {
	datas := service.repo.ListComment(req)
	dataset := []model.ListCommentResponse{}
	for _, data := range datas {
		dataset = append(dataset, model.ListCommentResponse{
			ID:        data.ID,
			PostID:    data.PostID,
			ParentID:  data.ParentID,
			Content:   data.Content,
			CommentBy: data.CommentByID,
		})
	}

	relations := make(map[uint][]model.ListCommentResponse)

	for _, relation := range dataset {
		child, parent := relation, relation.ParentID
		relations[parent] = append(relations[parent], child)
	}

	comments := buildComment(relations[0], relations)

	return comments
}
