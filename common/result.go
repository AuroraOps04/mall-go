package common

import (
	"github.com/gin-gonic/gin"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, Result{
		Code:    200,
		Message: "ok",
		Data:    data,
	})
}

func Error(ctx *gin.Context, msg string) {
	ctx.JSON(200, Result{
		Code:    400,
		Message: msg,
		Data:    nil,
	})
}

type PageResult struct {
	List      interface{} `json:"list,omitempty"`
	PageNum   string      `json:"pageNum,omitempty"`
	PageSize  string      `json:"pageSize,omitempty"`
	Total     int64       `json:"total,omitempty"`
	TotalPage string      `json:"totalPage,omitempty"`
}

func Page(ctx *gin.Context, data interface{}, count int64) {
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": PageResult{
			List:     data,
			Total:    count,
			PageNum:  ctx.Query("pageNum"),
			PageSize: ctx.Query("pageSize"),
		},
		"message": "ok",
	})
}
