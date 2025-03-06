package models

type Student struct {
	ID       int      `json:"id"`
	Number   string   `json:"number"`
	Password string   `json:"password"`
	ClassId  int      `json:"class_id"`
	Name     string   `json:"name"`
	Lesson   []Lesson `gorm:"many2many:lesson_student"`
}

func (Student) TableName() string {
	return "student"
}
