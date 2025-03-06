package models

type LessonStudent struct {
	StudentId string `json:"student_id"`
	LessonId  string `json:"lesson_id"`
}

func (LessonStudent) TableName() string {
	return "lesson_student"
}
