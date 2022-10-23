package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	SectionID   uint
	Title       string
	Description string
	Slug        string
	Order       int
	HTML        *string
	MarkDown    *string
	Published   bool `gorm:"default:false"`
	//Comments
}
