package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pathak107/coderahi-learn/pkg/dto"
	"github.com/pathak107/coderahi-learn/pkg/models"
	"github.com/pathak107/coderahi-learn/pkg/post"
)

var (
	post_cache      = "posts"
	post_ID_cache   = "post-ID-"
	post_Slug_cache = "post-Slug-"
)

func (h *Handler) FindAllPosts(ctx *gin.Context) {
	cacheResp := h.cache.Get(post_cache)
	if cacheResp.Found {
		logger.Printf("Cache hit for %v\n", post_cache)
		posts := cacheResp.Data.([]models.Post)
		ctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"posts": posts,
			},
		})
		return
	}

	logger.Printf("Cache miss for %v, querying from db\n", post_cache)
	posts, err := post.FindAllPosts(h.db)
	if err != nil {
		ctx.Error(err)
		return
	}

	h.cache.Put(post_cache, posts)

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"posts": posts,
		},
	})
}

func (h *Handler) FindPostByID(ctx *gin.Context) {
	postID := ctx.Param("post_id")

	cacheResp := h.cache.Get(post_ID_cache + postID)
	if cacheResp.Found {
		logger.Printf("Cache hit for %v\n", post_ID_cache+postID)
		post := cacheResp.Data.(models.Post)
		ctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"post": post,
			},
		})
		return
	}

	logger.Printf("Cache miss for %v, querying from db\n", post_ID_cache+postID)
	post, err := post.FindPostByID(h.db, postID)
	if err != nil {
		ctx.Error(err)
		return
	}

	h.cache.Put(post_ID_cache+postID, post)

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"post": post,
		},
	})
}

func (h *Handler) FindPostBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	cacheResp := h.cache.Get(post_Slug_cache + slug)
	if cacheResp.Found {
		logger.Printf("Cache hit for %v\n", post_Slug_cache+slug)
		post := cacheResp.Data.(models.Post)
		ctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"post": post,
			},
		})
		return
	}

	logger.Printf("Cache miss for %v, querying from db\n", post_Slug_cache+slug)
	post, err := post.FindPostBySlug(h.db, slug)
	if err != nil {
		ctx.Error(err)
		return
	}

	h.cache.Put(post_Slug_cache+slug, post)

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

	defer h.cache.Remove(post_cache)

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

	defer h.cache.Purge()

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

	defer h.cache.Purge()

	ctx.JSON(http.StatusOK, gin.H{
		"data": "post saved succesfully",
	})
}

func (h *Handler) ChangePostOrder(ctx *gin.Context) {
	var changeOrderDTO dto.ChangeOrderPostDTO
	if err := ctx.ShouldBind(&changeOrderDTO); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := post.ChangeOrderOfPost(h.db, &changeOrderDTO)
	if err != nil {
		ctx.Error(err)
		return
	}

	defer h.cache.Purge()

	ctx.JSON(http.StatusOK, gin.H{
		"data": "changed post order successfully",
	})
}
