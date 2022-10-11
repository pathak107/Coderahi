package post

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrPostNotFound = errors.New("post not found")
)

type Post struct {
	gorm.Model
	SectionID   uint
	Title       string
	Description string
	Slug        string
	HTMLBody    *string
	MarkDown    *string
	ImageUrl    *string
	//Comments
}
