package entity

import "time"

type User struct {
	ID           uint      `gorm:"column:id"`
	UserName     string    `gorm:"column:user_name"`
	Phone        string    `gorm:"column:phone"`
	Email        string    `gorm:"column:email"`
	PasswordHash string    `gorm:"password_hash"`
	RegisteredAt time.Time `gorm:"column:registered_at"`
	LastLogin    time.Time `gorm:"column:last_login"`
	Profile      string    `gorm:"column:profile"`
}

func (User) TableName() string {
	return "user"
}
