package models

import "time"

// type Navs struct {
// 	ID      int       `json:"id"`
// 	Title   string    `json:"title"`
// 	Url     string    `json:"url"`
// 	Status  int       `json:"status"`
// 	Sort    int       `json:"sort"`
// 	AddTime time.Time `json:"add_time"`
// }

type Navs struct {
	ID      int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title   string    `gorm:"size:100" json:"title"` // 限制最大长度为 100
	Url     string    `gorm:"size:2083" json:"url"`  // URL 的最大长度建议 2083
	Status  int       `json:"status"`
	Sort    int       `json:"sort"`
	AddTime time.Time `json:"add_time"`
}

// func (n Navs) TableName() string {
// 	return "navs"
// }
