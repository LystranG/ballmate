package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ballmate/config"
	"ballmate/middleware"
	"ballmate/model"
	"ballmate/router"
	"ballmate/service"
)

func main() {
	// 创建必要目录
	os.MkdirAll("./data", 0755)
	os.MkdirAll("./uploads", 0755)

	// 初始化数据库
	db, err := gorm.Open(sqlite.Open(config.DBPath), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// SQLite 并发写入限制
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("获取数据库实例失败:", err)
	}
	sqlDB.SetMaxOpenConns(1)

	// 自动迁移
	db.AutoMigrate(&model.User{})

	// 注入数据库实例到 service 层
	service.DB = db

	// 初始化 Gin
	r := gin.Default()

	// 注册 CORS 中间件
	r.Use(middleware.CorsMiddleware())

	// 注册路由
	router.SetupRouter(r)

	// 启动服务
	log.Println("服务启动在 :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("服务启动失败:", err)
	}
}
