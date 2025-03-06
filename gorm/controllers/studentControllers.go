package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leebrouse/gorm/models"
	"gorm.io/gorm"
)

type IStudentControllers interface {
	IBaseControllers
	Index(c *gin.Context)
}

type StudentControllers struct {
	BaseControllers
}

func NewStudentControllers() IStudentControllers {
	return &StudentControllers{}
}

func (student *StudentControllers) Index(c *gin.Context) {
	// studentLists := []models.Student{}
	// models.DB.Preload("Lesson").Where("name=?", "张三").Find(&studentLists)

	// c.JSON(http.StatusOK, gin.H{
	// 	"result": studentLists,
	// })

	// lessonLists := []models.Lesson{}
	// models.DB.Preload("Student", "id not in (?,?)", 1, 2).Limit(2).Find(&lessonLists)
	// c.JSON(http.StatusOK, gin.H{
	// 	"result": lessonLists,
	// })

	lessonLists := []models.Lesson{}
	models.DB.Preload("Student", func(db *gorm.DB) *gorm.DB {
		return models.DB.Where("id>3").Order("Student.id desc")
	}).Limit(2).Find(&lessonLists)
	c.JSON(http.StatusOK, gin.H{
		"result": lessonLists,
	})
}
