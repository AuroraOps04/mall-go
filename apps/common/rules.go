package common

import "github.com/gin-gonic/gin"

func Route(engine *gin.Engine) {
	engine.POST("/admin/minio/upload", UploaderController.Upload)
}
