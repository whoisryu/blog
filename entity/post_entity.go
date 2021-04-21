package entity

import "time"

type Post struct {
	ID          uint       `gorm:"column:id" json:"id"`
	AuthorId    uint       `gorm:"column:author_id" json:"author_id"`
	Author      User       `gorm:"foreignKey:ID;references:author_id" json:"-"`
	Title       string     `gorm:"column:title" json:"title"`
	Slug        string     `gorm:"columm:slug" json:"slug"`
	IsPublished bool       `gorm:"column:is_published" json:"is_published"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at"`
	PublishedAt time.Time  `gorm:"column:published_at" json:"published_at"`
	Content     string     `gorm:"column:content" json:"content"`
	Views       int        `grom:"columm:views" json:"-"`
	Categories  []Category `gorm:"many2many:post_category" json:"categories"`
}

func (Post) TableName() string {
	return "post"
}
