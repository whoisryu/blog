package entity

type PostCategory struct {
	ID uint `gorm:"columm:id"`
}

func (PostCategory) TableName() string {
	return "post_category"
}
