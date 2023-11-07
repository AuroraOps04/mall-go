package common

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PreparePage(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	var page = struct {
		PageNum  int `json:"pageNum"`
		PageSize int `json:"pageSize"`
	}{}
	err := ctx.ShouldBindQuery(&page)
	var limit, offset int
	if err != nil {
		limit = 10
		offset = 0
	} else {
		limit = page.PageSize
		pageNum := page.PageNum - 1
		offset = pageNum * limit
	}
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit).Offset(offset)
	}
}
