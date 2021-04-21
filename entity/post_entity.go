package entity

import "time"

type Post struct {
	ID          uint      `gorm:"column:id"`
	AuthorId    uint      `gorm:"column:author_id"`
	Author      User      `gorm:"foreignKey:ID;references:author_id"`
	Title       string    `gorm:"column:title"`
	Slug        string    `gorm:"columm:slug"`
	IsPublished bool      `gorm:"column:is_published"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	PublishedAt time.Time `gorm:"column:published_at"`
	Content     string    `gorm:"column:content"`
	Views       int       `grom:"columm:views"`
}

func (Post) TableName() string {
	return "post"
}
