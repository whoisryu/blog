package entity

import "time"

type User struct {
	ID           uint      `gorm:"column:id" json:"id"`
	UserName     string    `gorm:"column:user_name" json:"username"`
	Phone        string    `gorm:"column:phone" json:"phone"`
	Email        string    `gorm:"column:email" json:"email"`
	PasswordHash string    `gorm:"password_hash" json:"-"`
	RegisteredAt time.Time `gorm:"column:registered_at" json:"register_at"`
	LastLogin    time.Time `gorm:"column:last_login" json:"last_login"`
	Profile      string    `gorm:"column:profile" json:"profile"`
}

func (User) TableName() string {
	return "user"
}
