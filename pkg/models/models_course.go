package models

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Title        string
	DescHTML     *string
	DescJson     *string
	DescMarkdown *string
	ImageURL     *string
	DescShort    string
	Slug         string
	Cost         int
	Sections     []Section `gorm:"constraint:OnDelete:CASCADE;"`
	ExpectedTime int
	Likes        int
	Views        int
	Categories   []Category `gorm:"many2many:course_categories;"`
	Published    bool       `gorm:"default:false"`
	//Comments
}

type Section struct {
	gorm.Model
	CourseID     uint
	Title        string
	Description  string
	ExpectedTime int //in minutes
	Order        int
	Published    bool   `gorm:"default:false"`
	Posts        []Post `gorm:"constraint:OnDelete:CASCADE;"`
}
