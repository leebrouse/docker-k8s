package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leebrouse/gorm/models"
)

// interface
type IuserControllers interface {
	Index(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

// struct
type UserControllers struct{}

// function
func NewUserController() IuserControllers {
	return &UserControllers{}
}

// index
func (users *UserControllers) Index(c *gin.Context) {

	userList := []models.Users{}
	models.DB.Find(&userList)

	//age>20 users
	// userList := []models.Users{}
	// models.DB.Where("age>?", 25).Find(&userList)

	c.JSON(http.StatusOK, gin.H{
		"message": "index user",
		"resulet": userList,
	})
}

// add
func (users *UserControllers) Add(c *gin.Context) {

	user := models.Users{
		UserName: "Leebrouse",
		Age:      19,
		Email:    "leebrouse@gmail.com",
	}
	models.DB.Select("UserName", "Age", "Email").Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "add user",
		"success": true,
	})
}

// update
func (users *UserControllers) Update(c *gin.Context) {

	// user := models.Users{}
	// models.DB.First(&user)

	// user.Age = 20
	// user.Email = "2313186065@qq.com"
	// models.DB.Select("Id", "UserName", "Age").Save(&models.Users{
	// 	Id:       6,
	// 	UserName: "helloworld",
	// 	Age:      30,
	// })

	models.DB.Model(&models.Users{}).Where("id=?", 30).Update("username", "SongYuTang")

	c.JSON(http.StatusOK, gin.H{
		"message": "update user",
		"success": true,
	})
}

// delete
func (users *UserControllers) Delete(c *gin.Context) {
	user := models.Users{}

	// models.DB.Where("username=?", "Leebrouse").Delete(&user)
	models.DB.Delete(&user, []int{1, 2, 3})

	c.JSON(http.StatusOK, gin.H{
		"message": "delete user",
		"success": true,
	})
}
