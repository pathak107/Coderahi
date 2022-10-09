package handler

import (
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
