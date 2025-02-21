// go get github.com/spf13/viper
package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	// // 设置配置文件名（不带扩展名）
	// viper.SetConfigName("viper")
	// // 设置配置文件类型（这里是 YAML）
	// viper.SetConfigType("yaml")
	// // 设置配置文件路径（当前目录）
	// viper.AddConfigPath(".")

	viper.SetConfigFile("./viper.yaml")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置失败: %v", err)
	}

	// 获取配置值
	fmt.Println("app_name:", viper.GetString("app_name"))
	fmt.Println("port:", viper.GetInt("port"))
	fmt.Println("database user:", viper.GetString("database.user"))
	fmt.Println("database password:", viper.GetString("database.password"))

}
