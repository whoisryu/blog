package entity

type Category struct {
	ID       uint   `gorm:"column:id" json:"id"`
	ParentID uint   `gorm:"column:parent_id" json:"parent_id"`
	Title    string `gorm:"column:title" json:"title"`
	Slug     string `gorm:"column:slug" json:"slug"`
	Posts    []Post `gorm:"many2many:post_category" json:"posts"`
}

func (Category) TableName() string {
	return "category"
}
