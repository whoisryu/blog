package model

import "time"

type ListPostRequest struct {
	SortBy uint   `json:"sort_by"` // 0 terbaru 1 terpopuler
	Q      string `json:"q"`
}

type ListPostResponse struct {
	ID         uint   `json:"id"`
	AuthorName string `gorm:"column:user_name" json:"author_name"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Content    string `json:"content"`
}

type CreatePostRequest struct {
	Title       string `json:"title"`
	IsPublished bool   `json:"is_published"`
	Content     string `json:"content"`
	Categories  []uint `json:"categories"`
	AuthorId    uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt time.Time
}

type UpdatePostRequest struct {
	ID          string
	Title       string `json:"title"`
	IsPublished bool   `json:"is_published"`
	Content     string `json:"content"`
	Categories  []uint `json:"categories"`
	AuthorId    uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt time.Time
}
