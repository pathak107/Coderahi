package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pathak107/coderahi-learn/pkg/category"
)

func (h *Handler) FindAllCategories(ctx *gin.Context) {
	cats, err := category.FindAllCategories(h.db)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"categories": cats,
		},
	})
}
