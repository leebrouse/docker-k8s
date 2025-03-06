package models

type Article struct {
	ID          int
	Title       string
	CateId      int
	State       int
	ArticleCate ArticleCate `gorm:"foreignKey:CateId;references:ID"`
}

func (Article) TableName() string {
	return "article"
}
