package services

import (
	"fmt"
	"os"

	"github.com/pathak107/coderahi-learn/pkg/course"
	"github.com/pathak107/coderahi-learn/pkg/post"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseService() (*gorm.DB, error) {
	dbURL := os.Getenv("DB_URL")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/postgres", dbUsername, dbPassword, dbURL, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&post.Post{}, &course.Course{}, &course.Section{}, &course.Subsection{})
	return db, err
}
