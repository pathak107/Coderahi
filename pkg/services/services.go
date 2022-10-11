package services

import (
	"os"

	"github.com/pathak107/coderahi-learn/pkg/course"
	"github.com/pathak107/coderahi-learn/pkg/post"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseService() (*gorm.DB, error) {
	dbURL := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	db.AutoMigrate(&post.Post{}, &course.Course{}, &course.Section{})
	return db, err
}
