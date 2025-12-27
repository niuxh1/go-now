package articles

import (
	"bufio"
	"fmt"
	database "go-now-blog-backend/DataBase"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ImportPostsFromDir() {
	candidateDirs := []string{"posts", "backend/posts", "../backend/posts"}
	var postsDir string
	
	for _, dir := range candidateDirs {
		if _, err := os.Stat(dir); err == nil {
			postsDir = dir
			break
		}
	}

	if postsDir == "" {
		cwd, _ := os.Getwd()
		fmt.Printf("Warning: Could not find 'posts' directory. Current working directory: %s\n", cwd)
		return
	}

	absDir, _ := filepath.Abs(postsDir)
	fmt.Printf("Scanning articles from: %s\n", absDir)

	files, err := os.ReadDir(postsDir)
	if err != nil {
		fmt.Printf("Warning: Could not read posts directory: %v\n", err)
		return
	}

	db := database.GetDB()
	count := 0

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".md" {
			continue
		}

		path := filepath.Join(postsDir, file.Name())
		article, err := parseMarkdownFile(path)
		if err != nil {
			fmt.Printf("Skipping %s: %v\n", file.Name(), err)
			continue
		}

		var existing database.Article
		var found bool

		if article.ID != 0 {
			if err := db.Unscoped().First(&existing, article.ID).Error; err == nil {
				found = true
			}
		} else {
			if err := db.Unscoped().Where("title = ?", article.Title).First(&existing).Error; err == nil {
				found = true
			}
		}

		if !found {
			if article.Date.IsZero() {
				article.Date = time.Now()
			}
			if err := db.Create(article).Error; err == nil {
				fmt.Printf("Imported: %s (ID: %d)\n", article.Title, article.ID)
				count++
			} else {
				fmt.Printf("Failed to import %s: %v\n", article.Title, err)
			}
		} else {
			existing.Title = article.Title
			existing.Content = article.Content
			existing.Summary = article.Summary
			existing.Category = article.Category
			
			existing.DeletedAt = gorm.DeletedAt{} 

			if !article.Date.IsZero() {
				existing.Date = article.Date
			}
			
			if err := db.Save(&existing).Error; err == nil {
				fmt.Printf("Updated: %s (ID: %d)\n", existing.Title, existing.ID)
				count++
			} else {
				fmt.Printf("Failed to update %s: %v\n", existing.Title, err)
			}
		}
	}

	if count > 0 {
		fmt.Printf("Successfully imported %d articles from %s\n", count, postsDir)
	}
}

func TriggerImport(c *gin.Context) {
	ImportPostsFromDir()
	c.JSON(http.StatusOK, gin.H{"message": "Import process triggered. Check server logs for details."})
}

func parseMarkdownFile(path string) (*database.Article, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var frontMatter []string
	var contentBuilder strings.Builder
	inFrontMatter := false
	lineNum := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		if line == "---" {
			if lineNum == 1 {
				inFrontMatter = true
				continue
			} else if inFrontMatter {
				inFrontMatter = false
				continue
			}
		}

		if inFrontMatter {
			frontMatter = append(frontMatter, line)
		} else {
			contentBuilder.WriteString(line + "\n")
		}
	}

	article := &database.Article{
		Content: strings.TrimSpace(contentBuilder.String()),
		Views:   0,
	}

	for _, line := range frontMatter {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		switch key {
		case "id":
			if id, err := strconv.Atoi(val); err == nil {
				article.ID = uint(id)
			}
		case "title":
			article.Title = val
		case "date":
			if t, err := time.Parse("2006-01-02", val); err == nil {
				article.Date = t
			}
		case "category":
			article.Category = val
		case "summary":
			article.Summary = val
		}
	}

	if article.Title == "" {
		filename := filepath.Base(path)
		article.Title = strings.TrimSuffix(filename, filepath.Ext(filename))
	}

	if article.Category == "" {
		article.Category = "Tech"
	}

	if article.Summary == "" && len(article.Content) > 0 {
		end := 200
		if len(article.Content) < 200 {
			end = len(article.Content)
		}
		article.Summary = article.Content[:end] + "..."
	}

	return article, nil
}
