package auth

import (
	database "go-now-blog-backend/DataBase"
	"go-now-blog-backend/go-jwt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()
	var user database.User
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}
	token, err := jwt.GetToken(int64(user.ID), user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
