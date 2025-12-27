package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	// 优先从环境变量获取 DSN，如果没有则使用默认本地配置
	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = "root:197911@tcp(127.0.0.1:3306)/go_now_blog?charset=utf8mb4&parseTime=True&loc=Local"
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Database connection established successfully.")
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
