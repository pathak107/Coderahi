package services

import (
	"os"

	"github.com/pathak107/coderahi-learn/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseService() (*gorm.DB, error) {
	dbURL := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	db.AutoMigrate(&models.Post{}, &models.Course{}, &models.Section{})
	return db, err
}
