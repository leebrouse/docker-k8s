package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//add routers
	return r
}
func main() {
	r := NewRouter()
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Can`t connect with the mysql,due to:" + err.Error())
		return
	}

	fmt.Println(db)
	r.Run()
}
