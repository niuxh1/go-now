package comment

import (
	database "go-now-blog-backend/DataBase"
	"net/http"
	"strconv"
	"time"

	"strings"

	"github.com/gin-gonic/gin"
	"go-now-blog-backend/go-jwt"
)

type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type CommentResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	User      UserInfo  `json:"user"`
}

type AddCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

// GetComments 获取文章评论
func GetComments(c *gin.Context) {
	articleID := c.Param("id")
	db := database.GetDB()
	var comments []database.Comment

	if err := db.Where("article_id = ?", articleID).Preload("User").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	var response []CommentResponse
	for _, comment := range comments {
		response = append(response, CommentResponse{
			ID:        comment.ID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt,
			User: UserInfo{
				ID:       comment.User.ID,
				Username: comment.User.Username,
			},
		})
	}

	// Return empty list instead of null
	if response == nil {
		response = []CommentResponse{}
	}

	c.JSON(http.StatusOK, response)
}

// AddComment 添加评论
func AddComment(c *gin.Context) {
	articleIDStr := c.Param("id")
	articleID, err := strconv.Atoi(articleIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
		return
	}

	var req AddCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()

	// Try to get user from token
	var userID uint
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && parts[0] == "Bearer" {
			claims, err := jwt.ParseToken(parts[1])
			if err == nil {
				userID = uint(claims.UserID)
			}
		}
	}

	// If no valid user ID, try to find Guest user
	if userID == 0 {
		var guest database.User
		if err := db.Where("username = ?", "Guest").First(&guest).Error; err == nil {
			userID = guest.ID
		} else {
			// Fallback if Guest user is missing (should be created by InitGuest)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Guest user not configured"})
			return
		}
	}

	// Verify article exists
	var article database.Article
	if err := db.First(&article, articleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	comment := database.Comment{
		Content:   req.Content,
		ArticleID: uint(articleID),
		UserID:    userID,
	}

	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	// Preload user for response
	db.Preload("User").First(&comment, comment.ID)

	c.JSON(http.StatusCreated, CommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		User: UserInfo{
			ID:       comment.User.ID,
			Username: comment.User.Username,
		},
	})
}