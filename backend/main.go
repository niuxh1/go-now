package main

import (
	"fmt"
	database "go-now-blog-backend/DataBase"
	"go-now-blog-backend/articles"
	"go-now-blog-backend/auth"
	"go-now-blog-backend/comment"
	jwt "go-now-blog-backend/go-jwt"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 连接数据库
	db := database.Connect()

	// 2. 自动迁移表结构
	db.AutoMigrate(
		&database.User{},
		&database.Article{},
		&database.Comment{},
		&database.Tag{},
		&database.ArticleTag{},
		&database.SiteSetting{},
	)

	// 3. 初始化管理员和访客账户
	auth.InitAdmin()
	auth.InitGuest()

	// 4. 导入 posts/ 目录下的 Markdown 文章
	articles.ImportPostsFromDir()

	// 5. 配置路由
	r := gin.Default()

	// CORS 中间件
	r.Use(corsMiddleware())

	api := r.Group("/api")
	{
		// 认证模块
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/register", auth.Register)
			authGroup.POST("/login", auth.Login)
			authGroup.POST("/reset-password", auth.ResetPassword)
		}

		// 文章模块（公开）
		api.GET("/articles", articles.ShowAllArticle)
		api.GET("/articles/:id", articles.ShowArticleByID)

		// 文章模块（需要认证）
		api.POST("/articles", jwt.JwtMiddleWare(), articles.CreateArticle)
		api.DELETE("/articles/:id", jwt.JwtMiddleWare(), articles.DeleteArticle)

		// 评论模块
		api.GET("/articles/:id/comments", comment.GetComments)
		api.POST("/articles/:id/comments", comment.AddComment)

		// 标签
		api.GET("/tags", getTags)

		// 站点设置
		api.GET("/settings", getSettings)
		api.PUT("/settings", jwt.JwtMiddleWare(), updateSettings)
	}

	fmt.Println("Server starting on :8080")
	r.Run(":8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func getTags(c *gin.Context) {
	db := database.GetDB()
	var tags []database.Tag
	db.Find(&tags)
	c.JSON(200, tags)
}

func getSettings(c *gin.Context) {
	db := database.GetDB()
	var settings []database.SiteSetting
	db.Find(&settings)
	c.JSON(200, settings)
}

func updateSettings(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(403, gin.H{"error": "Only admin can update settings"})
		return
	}

	var settings []database.SiteSetting
	if err := c.ShouldBindJSON(&settings); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()
	for _, s := range settings {
		db.Save(&s)
	}
	c.JSON(200, gin.H{"message": "Settings updated"})
}
