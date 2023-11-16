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

// GetById godoc
//
//	@Param		id	path		int	true	"brand id"	default(1)
//	@Success	200	{object}	common.Result{data=pms.Brand}
//	@Tags		PmsBrandController
//	@Router		/brand/{id} [get]
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

// Create godoc
//
//	@Tags	PmsBrandController
//	@Param	brand	body	pms.Brand	true	"add brand"
//	@Accept	json
//	@Router	/brand/create [post]
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

// Page godoc
//
//	@Param		keyword		query	string	false	"keyword"
//	@Param		showStatus	query	string	false	"keyword"	Enums(0, 1)
//	@Tags		PmsBrandController
//	@Success	200	{object}	common.PageResult{data=[]pms.Brand}
//	@Router		/brand/list [get]
func (c *brandController) Page(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	showStatus := ctx.Query("showStatus")

	var brands []*pms.Brand
	tx := global.Db.Model(&pms.Brand{})
	if strings.TrimSpace(keyword) != "" {
		likeStr := "%" + keyword + "%"
		tx.Where("name like ?  ", likeStr)
	}
	if strings.TrimSpace(showStatus) != "" {
		tx.Where("show_status = ?", showStatus)
	}
	var count int64
	tx.Count(&count)
	tx.Scopes(common.PreparePage(ctx)).Find(&brands)
	common.Page(ctx, brands, count)
}

// DeleteById godoc
//
//	@Param		id	path	int	true	"brand id"
//	@Tags		PmsBrandController
//	@Success	200	{object}	common.Result{data=int}
//	@Router		/brand/delete/{id} [get]
func (c *brandController) DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")
	tx := global.Db.Delete(&pms.Brand{}, id)
	if tx.Error != nil || tx.RowsAffected == 0 {
		common.Error(ctx, "delete failed")
		return
	}
	common.Success(ctx, tx.RowsAffected)
}

// UpdateShowStatusBatch godoc
//
//	@Param		showStatus	formData	string	true	"show status"			Enums(0,1)
//	@Param		ids			formData	string	true	"id array join as ,"	"1,2"
//	@Tags		PmsBrandController
//
//	@Success	200	{object}	common.Result{data=int}
//
//	@Router		/brand/update/showStatus [post]
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

// UpdateFactoryStatusBatch godoc
//
//	@Summary	update brand show status batch
//
//	@Param		factoryStatus	formData	string	true	"factory status"
//	@Param		ids				formData	string	true	"id array join as ,"	"1,2"
//	@Tags		PmsBrandController
//
//	@Success	200	{object}	common.Result{data=int}
//
//	@Router		/brand/update/factoryStatus [post]
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

// Update godoc
//
//	@Summary	Update Brand by id
//	@Param		id		path	int			true	"id"
//	@Param		brand	body	pms.Brand	true	"update brand"
//	@Tags		PmsBrandController
//
//	@Success	200	{object}	common.Result{data=int}
//
//	@Router		/brand/update/{id} [post]
func (c *brandController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var b pms.Brand
	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		common.Error(ctx, "params error")
		return
	}
	tx := global.Db.Where("id = ?", id).Save(&b)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Success(ctx, tx.RowsAffected)
}

// ListAll godoc
//
//	@Summary	list all brand
//	@Tags		PmsBrandController
//	@Success	200	{object}	common.Result{data=[]pms.Brand}
//	@Router		/brand/listAll [get]
func (c *brandController) ListAll(ctx *gin.Context) {
	var b []*pms.Brand
	global.Db.Model(&pms.Brand{}).Find(&b)
	common.Success(ctx, b)
}

// DeleteBatch godoc
//
//	@Summary	delete brand batch by ids
//	@Params		ids formData []string true
//	@Tags		PmsBrandController
//	@Success	200	{object}	common.Result{data=[]pms.Brand}
//	@Router		/brand/delete/batch [post]
func (c *brandController) DeleteBatch(ctx *gin.Context) {
	ids := ctx.PostFormArray("ids")
	tx := global.Db.Delete(&pms.Brand{}, ids)
	if tx.Error != nil || tx.RowsAffected == 0 {
		common.Error(ctx, "remove brand batch failed")
		return
	}
	common.Success(ctx, tx.RowsAffected)

}
