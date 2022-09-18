package handler

import (
	"encoding/json"
	"strconv"

	"github.com/pathak107/coderahi-learn/pkg/post"
	"github.com/pathak107/coderahi-learn/pkg/utils"
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

func (h *Handler) FindAllPosts() (string, error) {
	posts, err := post.FindAllPosts(h.db)
	if err != nil {
		return "", err
	}
	return NewDataRespHandler(posts)
}

func (h *Handler) FindPostByID(postID string) (string, error) {
	pID, _ := strconv.Atoi(postID)
	post, err := post.FindPostByID(h.db, uint(pID))
	if err != nil {
		return "", err
	}
	return NewDataRespHandler(post)
}

func (h *Handler) CreatePost(body string) (string, error) {
	var req post.CreatePostDTO
	err := json.Unmarshal([]byte(body), &req)
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	postID, err := post.CreatePost(h.db, &req)
	if err != nil {
		return "", err
	}
	return NewDataRespHandler(postID)
}

func (h *Handler) DeletePost(postID string) (string, error) {
	pID, _ := strconv.Atoi(postID)
	err := post.DeletePostByID(h.db, uint(pID))
	if err != nil {
		return "", err
	}
	return NewSuccessDeletionRespHandler()
}

func (h *Handler) EditPost(body string, postID string) (string, error) {
	var req post.EditPostDTO
	pID, _ := strconv.Atoi(postID)
	err := json.Unmarshal([]byte(body), &req)
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	req.PostID = uint(pID)
	err = post.EditPost(h.db, &req)
	if err != nil {
		return "", err
	}
	return NewSuccessEditRespHandler()
}
