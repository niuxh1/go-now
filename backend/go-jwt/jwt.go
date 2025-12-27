package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const Salt = "go-now-blog"

type JwtClaims struct {

	//   "user_id": 1,       // (Number) 用户唯一标识
	//   "username": "admin",// (String) 用户名
	//   "role": "admin",    // (String) 角色权限标识，必须为 "admin" 或 "user"
	//   "exp": 1735689600   // (Number) 过期时间戳
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GetToken(user_id int64, username, role string) (string, error) {
	claims := JwtClaims{
		UserID:   user_id,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			// ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 设置过期时间为24小时后
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Salt))
}

func ParseToken(token string) (*JwtClaims, error) {

	parsedToken, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(Salt), nil
	})

	if err != nil {
		return nil, err
	}

	if parsedToken.Valid {
		return parsedToken.Claims.(*JwtClaims), nil
	}
	return nil, errors.New("invalid token")
}

func JwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(401, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString
		claims, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set("claims", claims)
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}
