package models

type ArticleCate struct {
	ID      int
	Title   string
	State   int
	Article []Article `gorm:"foreignKey:CateId"`
}

func (ArticleCate) TableName() string {
	return "article_cate"
}
