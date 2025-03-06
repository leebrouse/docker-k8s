package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("Can`t connect with the mysql database,due to :" + err.Error())
		return
	}

	// 自动迁移
	err = DB.AutoMigrate(&Users{}, &Navs{}, &Bank{}, &Student{}, &Lesson{}, &LessonStudent{})
	if err != nil {
		log.Fatal("数据表迁移失败:", err)
	}

	fmt.Println("数据库连接成功，数据表已迁移")
}
