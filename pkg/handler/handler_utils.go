package handler

import (
	"github.com/pathak107/coderahi-learn/pkg/cache"
	"gorm.io/gorm"
)

type Handler struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewHandler(d *gorm.DB) *Handler {
	return &Handler{
		db:    d,
		cache: cache.NewCache(10),
	}
}
