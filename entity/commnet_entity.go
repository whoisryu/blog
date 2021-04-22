package entity

import "time"

type Comment struct {
	ID          uint      `gorm:"column:id" json:"id"`
	PostID      uint      `gorm:"column:post_id" json:"post_id"`
	ParentID    uint      `gorm:"column:parent_id" json:"parent_id"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	Content     string    `gorm:"column:content" json:"content"`
	CommentByID uint      `gorm:"column:comment_by" json:"comment_by"`
	CommentBy   User      `gorm:"association_foreignKey:ID" json:"-"`
}

func (Comment) TableName() string {
	return "post_comment"
}
