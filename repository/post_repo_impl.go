package repository

import (
	"blog/entity"
	"blog/exception"
	"blog/model"

	"gorm.io/gorm"
)

type postRepoImpl struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) PostRepo {
	return &postRepoImpl{db: db}
}

func (repo postRepoImpl) ListPost(req model.ListPostRequest) (response []model.ListPostResponse) {

	sortBy := "p.views"
	if req.SortBy == 1 {
		sortBy = "p.published_at"
	}

	query := repo.db.Table("post p").Select("p.id, p.title, p.slug, p.content, u.user_name").Joins("JOIN user u on u.id=p.author_id").Order(sortBy).Where("p.is_published=1")

	if req.Q != "" {
		query = query.Where("p.title like ?", "%"+req.Q+"%")
	}

	err := query.Find(&response).Error
	exception.PanicIfNeeded(err)

	return response
}

func (repo postRepoImpl) PostBySlug(slug string) (response model.ListPostResponse) {
	err := repo.db.Table("post p").Select("p.id, p.title, p.slug, p.content, u.user_name").Joins("JOIN user u on u.id=p.author_id").Where("p.is_published=1").Find(&response).Error

	exception.PanicIfNeeded(err)

	return response
}

func (repo postRepoImpl) CreatePost(post entity.Post) entity.Post {
	err := repo.db.Create(&post).Error
	exception.PanicIfNeeded(err)

	return post
}

func (repo postRepoImpl) PostByID(ID uint) (post entity.Post) {
	err := repo.db.Where("id=?", ID).First(&post).Error

	exception.PanicIfNeeded(err)

	return post
}

func (repo postRepoImpl) UpdatePost(post entity.Post) entity.Post {
	err := repo.db.Updates(&post).Error
	exception.PanicIfNeeded(err)

	return post
}

func (repo postRepoImpl) DeletePost(ID uint) {
	err := repo.db.Where("id = ?", ID).Delete(&entity.Post{}).Error
	exception.PanicIfNeeded(err)
}

func (repo postRepoImpl) ListByCategory(req model.ListPostByCategoryRequest) (response []model.ListPostResponse) {
	err := repo.db.Table("post").Joins("JOIN post_category ON post_category.post_id=post.id").Joins("JOIN category ON category.id=post_category.category_id").Joins("JOIN user on user.id=post.author_id").Where("category.slug=?", req.Slug).Select("post.id, user.user_name, post.title, post.slug, post.content").Find(&response).Error

	exception.PanicIfNeeded(err)

	return response
}
