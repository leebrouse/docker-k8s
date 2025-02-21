package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Database  DatabaseConfig  `mapstructure:"database"`
	Redis     RedisConfig     `mapstructure:"redis"`
	Server    ServerConfig    `mapstructure:"server"`
	App       AppConfig       `mapstructure:"app"`
	ShortCode ShortCodeConfig `mapstructure:"shortCode"`
}

type DatabaseConfig struct {
	// 添加数据库相关配置字段
	Driver       string `mapstructure:"driver"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	SSLMode      string `mapstructure:"ssl_mode"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

type RedisConfig struct {
	// 添加 Redis 相关配置字段
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type ServerConfig struct {
	// 添加服务器相关配置字段
	Addr         string        `mapstructure:"addr"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
}

type AppConfig struct {
	// 添加应用相关配置字段
	BaseURL         string        `mapstructure:"base_url"`
	DefaultDuration time.Duration `mapstructure:"default_duration"`
	CleanupInterval time.Duration `mapstructure:"cleanup_interval"`
}

type ShortCodeConfig struct {
	// 添加短代码相关配置字段
	Length int `mapstructure:"length"`
}

func LoadConfig(filename string) (*Config, error) {
	viper.SetConfigFile(filename)

	// 设置环境变量前缀为 URL_SHORTENER
	viper.SetEnvPrefix("URL_SHORTENER")
	// 将配置中的 `.` 替换为 `-`，这样可以支持用 `URL_SHORTENER_DATABASE-HOST` 形式的环境变量
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "-"))
	viper.AutomaticEnv() // 启用自动加载环境变量

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	// 将配置文件中的内容反序列化到 cfg 结构体中
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func main() {
	cfg, err := LoadConfig("./config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(*cfg)
}
