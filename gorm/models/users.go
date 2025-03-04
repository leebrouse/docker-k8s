package models

import "time"

// users struct
type Users struct {
	Id       int       `gorm:"column:id"`
	UserName string    `gorm:"column:username"`
	Age      int       `gorm:"column:age"`
	Email    string    `gorm:"column:email"`
	AddTime  time.Time `gorm:"column:add_time"`
}

func (Users) TableName() string {
	return "users"
}
