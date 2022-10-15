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
	CategoryID   uint
	//Comments
}

type Section struct {
	gorm.Model
	CourseID     uint
	Title        string
	Description  string
	ExpectedTime int //in minutes
	Order        int
	Posts        []Post `gorm:"constraint:OnDelete:CASCADE;"`
}

type Category struct {
	Name        string
	Slug        string
	Description string
	Courses     []Course
}
