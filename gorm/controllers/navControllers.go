package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leebrouse/gorm/models"
)

type INavControllers interface {
	IBaseControllers
	Index(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type NavControllers struct {
	BaseControllers
}

// function
func NewNavControllers() INavControllers {
	return &NavControllers{}
}

func (nav *NavControllers) Index(c *gin.Context) {
	// navs := []models.Navs{}
	//查找所有nav数据
	// models.DB.Find(&navs)

	//查找一个nav数据
	// navs := models.Navs{ID: 2}
	// models.DB.Find(&navs)
	// // models.DB.First(&navs)

	// c.JSON(http.StatusOK, gin.H{
	// 	"Navs":    navs,
	// 	"success": true,
	// })

	// 使用原生 sql 删除 user 表中的一条数据
	// query1 := `delete from users where id = ?`
	// models.DB.Exec(query1, 3)

	// 使用原生 sql 修改 user 表中的一条数据
	// query2 := `update users set username='leebrouse' where id=?`
	// models.DB.Exec(query2, 7)

	// 使用原生 sql 查询 uid=2 的数据
	// user := models.Users{}
	// query3 := `select * from users where id=?`
	// models.DB.Raw(query3, 7).Scan(&user)
	// fmt.Println(user)

	// 使用原生 查询 User 表中所有的数据
	// users := []models.Users{}
	// query4 := `select * from users`
	// models.DB.Raw(query4).Scan(&users)
	// fmt.Println(users)

	// 统计 user 表的数量
	// var number int
	// query5 := `select count(1) from users`
	// models.DB.Raw(query5).Scan(&number)
	// fmt.Println(number)

	articleList := []models.Article{}
	models.DB.Preload("ArticleCate").Find(&articleList)

	articleCateList := []models.ArticleCate{}
	models.DB.Preload("Article").Find(&articleCateList)

	c.JSON(http.StatusOK, gin.H{
		"Result":      articleList,
		"Result_cate": articleCateList,
		"success":     true,
	})

}

func (nav *NavControllers) Add(c *gin.Context) {

}

func (nav *NavControllers) Update(c *gin.Context) {

}

func (nav *NavControllers) Delete(c *gin.Context) {

}
