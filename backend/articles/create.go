package articles

import (
	"fmt"
	database "go-now-blog-backend/DataBase"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateArticleRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Date    string `json:"date"` // Optional, format: YYYY-MM-DD
}

func CreateArticle(c *gin.Context) {
	// 1. 验证权限 (必须是 admin)
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admin can create articles"})
		return
	}

	// 2. 绑定请求参数
	var req CreateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. 处理日期
	var publishDate time.Time
	var err error
	if req.Date != "" {
		publishDate, err = time.Parse("2006-01-02", req.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
			return
		}
	} else {
		publishDate = time.Now()
	}

	// 4. 创建文章
	article := database.Article{
		Title:    req.Title,
		Content:  req.Content,
		Date:     publishDate,
		Category: "Tech", // 默认分类，后续可扩展
		Summary:  generateSummary(req.Content),
	}

	db := database.GetDB()
	if err := db.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
		return
	}

	// 5. 同步保存到文件系统 (backend/posts/)
	go func(id uint, art database.Article) {
		filename := fmt.Sprintf("%d.md", id)
		path := filepath.Join("posts", filename)
		
		// 确保目录存在
		if _, err := os.Stat("posts"); os.IsNotExist(err) {
			os.Mkdir("posts", 0755)
		}

		// 构建带 Front Matter 的内容
		fileContent := fmt.Sprintf("---\nid: %d\ntitle: %s\ndate: %s\ncategory: %s\nsummary: %s\n---\n\n%s",
			art.ID,
			art.Title,
			art.Date.Format("2006-01-02"),
			art.Category,
			art.Summary,
			art.Content,
		)

		if err := os.WriteFile(path, []byte(fileContent), 0644); err != nil {
			fmt.Printf("Failed to save article file: %v\n", err)
		} else {
			fmt.Printf("Article saved to %s\n", path)
		}
	}(article.ID, article)

	c.JSON(http.StatusCreated, article)
}

func generateSummary(content string) string {
	if len(content) > 200 {
		return content[:200] + "..."
	}
	return content
}
