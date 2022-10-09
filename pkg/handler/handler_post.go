package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pathak107/coderahi-learn/pkg/dto"
	"github.com/pathak107/coderahi-learn/pkg/post"
)

func (h *Handler) FindAllPosts(ctx *gin.Context) {
	posts, err := post.FindAllPosts(h.db)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"posts": posts,
		},
	})
}

func (h *Handler) FindPostByID(ctx *gin.Context) {
	postID := ctx.Param("post_id")
	post, err := post.FindPostByID(h.db, postID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"post": post,
		},
	})
}

func (h *Handler) FindPostBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	post, err := post.FindPostBySlug(h.db, slug)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"post": post,
		},
	})
}

func (h *Handler) CreatePost(ctx *gin.Context) {
	var createPostDTO dto.CreatePostDTO
	if err := ctx.ShouldBind(&createPostDTO); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	postID, err := post.CreatePost(h.db, &createPostDTO)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"post_id": postID,
		},
	})
}

func (h *Handler) DeletePost(ctx *gin.Context) {
	postID := ctx.Param("post_id")
	err := post.DeletePostByID(h.db, postID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": "post deleted succesfully",
	})
}

func (h *Handler) EditPost(ctx *gin.Context) {
	postID := ctx.Param("post_id")
	var editPostDTO dto.EditPostDTO
	if err := ctx.ShouldBind(&editPostDTO); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := post.EditPost(h.db, &editPostDTO, postID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "post saved succesfully",
	})
}
