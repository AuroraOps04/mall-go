package pms

import (
	"github.com/AuroraOps04/mall-go/common"
	"github.com/AuroraOps04/mall-go/global"
	"github.com/AuroraOps04/mall-go/model/pms"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type brandController struct{}

var BrandController = &brandController{}

// GetById find by id
func (c *brandController) GetById(ctx *gin.Context) {
	b := pms.Brand{}
	ui := ctx.Param("id")
	id, err := strconv.ParseUint(ui, 10, 32)
	if err != nil {
		common.Error(ctx, "id must ge digit")
		return
	}
	b.ID = uint(id)
	err = global.Db.Model(&b).First(&b).Error
	if err != nil {
		common.Error(ctx, "not found")
		return
	}
	common.Success(ctx, b)
}

func (c *brandController) Create(ctx *gin.Context) {
	b := pms.Brand{}
	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		common.Error(ctx, "params fail")
		return
	}
	if strings.TrimSpace(b.FirstLetter) == "" {
		b.FirstLetter = string(([]rune(b.Name))[0:1])
	}
	tx := global.Db.Create(&b)
	err = tx.Error
	if err != nil {
		common.Error(ctx, "save error")
		return
	}
	common.Success(ctx, tx.RowsAffected)
}

func (c *brandController) Page(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	showStatus := ctx.Query("showStatus")

	var brands []*pms.Brand
	tx := global.Db.Model(&pms.Brand{}).Scopes(common.PreparePage(ctx))
	if strings.TrimSpace(keyword) != "" {
		likeStr := "%" + keyword + "%"
		tx.Where("name like ?  ", likeStr)
	}
	if strings.TrimSpace(showStatus) != "" {
		tx.Where("show_status = ?", showStatus)
	}
	var count int64
	tx.Count(&count)
	tx.Find(&brands)
	common.Page(ctx, brands, count)
}

func (c *brandController) DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")
	tx := global.Db.Delete(&pms.Brand{}, id)
	if tx.Error != nil || tx.RowsAffected == 0 {
		common.Error(ctx, "delete failed")
		return
	}
	common.Success(ctx, tx.RowsAffected)
}

func (c *brandController) UpdateShowStatusBatch(ctx *gin.Context) {

	showStatus := ctx.PostForm("showStatus")
	ids := ctx.PostForm("ids")
	idArray := strings.Split(strings.TrimSpace(ids), ",")
	tx := global.Db.Model(&pms.Brand{}).Where("id in ?", idArray).Update("show_status", showStatus)
	if tx.Error == nil {
		common.Success(ctx, tx.RowsAffected)
		return
	}
	common.Error(ctx, tx.Error.Error())
}

func (c *brandController) UpdateFactoryStatusBatch(ctx *gin.Context) {

	factoryStatus := ctx.Query("factoryStatus")
	ids := ctx.PostForm("ids")
	idArray := strings.Split(strings.TrimSpace(ids), ",")
	tx := global.Db.Model(&pms.Brand{}).Where("id in ?", idArray).Update("factory_status", factoryStatus)
	if tx.Error == nil {
		common.Success(ctx, tx.RowsAffected)
		return
	}
	common.Error(ctx, tx.Error.Error())
}

func (c *brandController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var b pms.Brand
	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		common.Error(ctx, "params error")
		return
	}
	tx := global.Db.Model(&pms.Brand{}).Where("id = ?", id).Updates(b)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Success(ctx, tx.RowsAffected)
}

func (c *brandController) ListAll(ctx *gin.Context) {
	var b []*pms.Brand
	global.Db.Model(&pms.Brand{}).Find(&b)
	common.Success(ctx, b)
}

func (c *brandController) DeleteBatch(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	tx := global.Db.Delete(&pms.Brand{}, ids)
	if tx.Error != nil || tx.RowsAffected == 0 {
		common.Error(ctx, "remove brand batch failed")
		return
	}
	common.Success(ctx, tx.RowsAffected)

}
