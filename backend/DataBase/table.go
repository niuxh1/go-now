package database

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Username  string         `gorm:"unique;not null;type:varchar(255)" json:"username"`
	Email     string         `gorm:"unique;not null;type:varchar(255)" json:"email"`
	Password  string         `gorm:"not null;type:varchar(255)" json:"-"` // Don't return password
	Role      string         `gorm:"default:'user';type:varchar(50)" json:"role"`
}

// Article 文章模型
type Article struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Title     string         `gorm:"not null;type:varchar(255)" json:"title"`
	Summary   string         `gorm:"type:varchar(1000)" json:"summary"` // 列表页摘要
	Content   string         `gorm:"type:longtext" json:"content"`      // 完整内容
	Date      time.Time      `gorm:"type:date" json:"date"`             // 发布日期
	Category  string         `gorm:"type:varchar(100)" json:"category"`
	Views     int            `gorm:"default:0" json:"views"`
}

// Comment 评论模型
type Comment struct {
	gorm.Model
	Content   string  `gorm:"not null;type:text"`
	ArticleID uint    `gorm:"not null"`
	Article   Article `gorm:"constraint:OnDelete:CASCADE;"` // 关联文章
	UserID    uint    `gorm:"not null"`
	User      User    `gorm:"constraint:OnDelete:CASCADE;"` // 关联用户
}

// Tag 标签模型
type Tag struct {
	gorm.Model
	Name string `gorm:"unique;not null;type:varchar(100)"`
}

// ArticleTag 文章-标签关联模型 (多对多)
type ArticleTag struct {
	ArticleID uint `gorm:"primaryKey"`
	TagID     uint `gorm:"primaryKey"`
}

// SiteSetting 站点配置模型 (Key-Value)
type SiteSetting struct {
	Key   string `gorm:"primaryKey;type:varchar(100)" json:"key"`
	Value string `gorm:"type:text" json:"value"` // JSON string
}
