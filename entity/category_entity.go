package entity

type Category struct {
	ID       uint   `gorm:"column:id"`
	ParentID uint   `gorm:"column:parent_id"`
	Title    string `gorm:"column:title"`
	Slug     string `gorm:"column:slug"`
	Posts    []Post `gorm:"many2many:post_category" json:"posts"`
}

func (Category) TableName() string {
	return "category"
}
