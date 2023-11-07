package common

import (
	"github.com/AuroraOps04/mall-go/common"
	"github.com/AuroraOps04/mall-go/dto"
	"github.com/gin-gonic/gin"
	"time"
)

type uploaderController struct {
}

var UploaderController = &uploaderController{}

const StaticDir = "./static/"

func (receiver *uploaderController) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		common.Error(ctx, "not found file that named `file`")
		return
	}
	day := time.Now().Format("20060102")
	path := day + "/" + file.Filename

	err = ctx.SaveUploadedFile(file, StaticDir+path)
	if err != nil {
		common.Error(ctx, "upload failed")
		return
	}
	var data dto.UploadDto
	data.Url = "http://localhost:4000/static/" + path
	data.Name = file.Filename
	common.Success(ctx, data)

}
