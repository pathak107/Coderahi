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

type CreatePostDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string
	IsBlogPost  bool `json:"is_blog_post"`
}

type EditPostDTO struct {
	PostID      uint
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string
	Body        string `json:"editorjs_body"`
	IsBlogPost  bool
}
