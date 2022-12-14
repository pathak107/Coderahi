package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name    string
	Slug    string
	Courses []Course `gorm:"many2many:course_categories;"`
}
