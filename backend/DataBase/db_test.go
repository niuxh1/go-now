package database

import (
	"testing"
)

// TestSchemaInitialization 用于验证数据库表结构是否已按照需求正确创建
// 运行测试命令: go test ./DataBase/... -v
func TestSchemaInitialization(t *testing.T) {
	// 1. 尝试连接数据库
	// 注意：确保你的 Docker MySQL 容器正在运行，且密码配置正确
	Connect()

	if DB == nil {
		t.Fatal("❌ Database connection failed: DB object is nil")
	}

	// 2. 验证所有必需的表是否存在
	expectedTables := []string{
		"users",
		"articles",
		"comments",
		"tags",
		"article_tags",
	}

	t.Log("Checking for existence of required tables...")
	for _, tableName := range expectedTables {
		if !DB.Migrator().HasTable(tableName) {
			t.Errorf("❌ Missing table: '%s' does not exist.", tableName)
		} else {
			t.Logf("✅ Table found: %s", tableName)
		}
	}

	// 3. 验证关键字段变更 (根据 API_REQUIREMENTS.md v2.0)
	// 检查 'users' 表是否包含 'email' 字段
	if DB.Migrator().HasTable("users") {
		if !DB.Migrator().HasColumn("users", "email") {
			t.Error("❌ Table 'users' is missing required column: 'email'")
		} else {
			t.Log("✅ Column found: users.email")
		}
	}

	// 检查 'articles' 表是否包含 'summary' 和 'category' 字段
	if DB.Migrator().HasTable("articles") {
		if !DB.Migrator().HasColumn("articles", "summary") {
			t.Error("❌ Table 'articles' is missing required column: 'summary'")
		} else {
			t.Log("✅ Column found: articles.summary")
		}

		if !DB.Migrator().HasColumn("articles", "category") {
			t.Error("❌ Table 'articles' is missing required column: 'category'")
		} else {
			t.Log("✅ Column found: articles.category")
		}
	}
}
