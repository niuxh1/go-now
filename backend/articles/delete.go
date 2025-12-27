package articles

import (
	"fmt"
	database "go-now-blog-backend/DataBase"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// DeleteArticle handles the deletion of an article
func DeleteArticle(c *gin.Context) {
	// 1. Check permissions (Admin only)
	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admin can delete articles"})
		return
	}

	// 2. Get Article ID from URL
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Article ID is required"})
		return
	}

	// 3. Find article to ensure it exists
	var article database.Article
	db := database.GetDB()
	if err := db.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// 4. Delete from Database
	if err := db.Delete(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete article from database"})
		return
	}

	// 5. Delete corresponding Markdown file
	// We run this synchronously to ensure consistency, unlike creation which was async
	filename := fmt.Sprintf("%d.md", article.ID)
	path := filepath.Join("posts", filename)
	
	if err := os.Remove(path); err != nil {
		// Even if file deletion fails, DB deletion succeeded, so we log it but don't fail the request
		// (Or we could treat it as a warning)
		fmt.Printf("Warning: Failed to delete article file %s: %v\n", path, err)
	} else {
		fmt.Printf("Deleted article file: %s\n", path)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}
