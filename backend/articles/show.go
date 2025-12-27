package articles

import (
	database "go-now-blog-backend/DataBase"

	"github.com/gin-gonic/gin"
)

func GetAllArticle(category string) []database.Article {
	db := database.GetDB()
	var articles []database.Article
	query := db.Order("date DESC") // Order by date descending by default
	if category != "" {
		query = query.Where("category = ?", category)
	}
	query.Find(&articles)
	return articles
}

func ShowAllArticle(c *gin.Context) {
	category := c.Query("category")
	articles := GetAllArticle(category)
	c.JSON(200, articles)
}

func ShowArticleByID(c *gin.Context) {
	id := c.Param("id")
	db := database.GetDB()
	var article database.Article
	if err := db.First(&article, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Article not found"})
		return
	}
	c.JSON(200, article)	
}