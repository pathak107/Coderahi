package handler

import (
	"strconv"

	"github.com/pathak107/coderahi-learn/pkg/post"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(d *gorm.DB) *Handler {
	return &Handler{
		db: d,
	}
}

func (h *Handler) FindAllPosts() ([]post.Post, error) {
	posts, err := post.FindAllPosts(h.db)
	if err != nil {
		return posts, err
	}
	return posts, nil
}

func (h *Handler) FindPostByID(postID string) (post.Post, error) {
	pID, _ := strconv.Atoi(postID)
	post, err := post.FindPostByID(h.db, uint(pID))
	if err != nil {
		return post, err
	}
	return post, nil
}
