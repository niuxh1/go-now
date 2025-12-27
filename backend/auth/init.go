package auth

import (
	"fmt"
	database "go-now-blog-backend/DataBase"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// InitAdmin 检查并创建/重置默认管理员
func InitAdmin() {
	db := database.GetDB()
	var user database.User
	
	// 生成强随机密码
	password := generateRandomPassword(16)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Failed to hash password: %v\n", err)
		return
	}

	// 查找是否有名为 "admin" 的用户
	result := db.Where("username = ?", "admin").First(&user)

	if result.Error == nil {
		// 用户存在：强制更新密码和角色
		user.Password = string(hashedPassword)
		user.Role = "admin"
		if err := db.Save(&user).Error; err != nil {
			fmt.Printf("Failed to update admin credentials: %v\n", err)
			return
		}
		fmt.Println("\n============================================")
		fmt.Println(" [Admin Account RESET] ")
	} else {
		// 用户不存在：创建新用户
		admin := database.User{
			Username: "admin",
			Email:    "admin@gonow.local",
			Password: string(hashedPassword),
			Role:     "admin",
		}
		if err := db.Create(&admin).Error; err != nil {
			fmt.Printf("Failed to create admin: %v\n", err)
			return
		}
		fmt.Println("\n============================================")
		fmt.Println(" [Admin Account CREATED] ")
	}

	fmt.Printf(" Username: admin\n")
	fmt.Printf(" Password: %s\n", password)
	fmt.Println(" -> Please save these credentials immediately!")
	fmt.Println("============================================\n")
}

// InitGuest Ensures a guest user exists for anonymous comments
func InitGuest() {
	db := database.GetDB()
	var user database.User
	if err := db.Where("username = ?", "Guest").First(&user).Error; err != nil {
		guest := database.User{
			Username: "Guest",
			Email:    "guest@gonow.local",
			Password: "guest_cannot_login", // Dummy password
			Role:     "user",
		}
		db.Create(&guest)
		fmt.Println(" [Guest Account CREATED] ")
	}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"

func generateRandomPassword(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
