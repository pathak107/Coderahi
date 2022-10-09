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
	Title       string
	Description string
	Slug        string
	HTMLBody    *string
	MarkDown    *string
	ImageUrl    *string
	Likes       int
	Views       int
	IsBlogPost  bool
	//Comments
}
