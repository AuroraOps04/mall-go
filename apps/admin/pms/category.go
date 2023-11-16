package pms

import (
	"github.com/AuroraOps04/mall-go/common"
	dto "github.com/AuroraOps04/mall-go/dto/pms"
	"github.com/AuroraOps04/mall-go/global"
	"github.com/AuroraOps04/mall-go/model/pms"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"strconv"
)

type productCategoryController struct {
}

var ProductCategoryController = &productCategoryController{}

// ListAllWithChildren
//
//	@Summary	list all category that parentId equals zero and preload children
//	@Tags		PmsProductCategoryController
//	@Success	200	{object}	common.Result{data=[]pms.ProductCategory}
//	@Router		/productCategory/list/withChildren [get]
func (p *productCategoryController) ListAllWithChildren(ctx *gin.Context) {
	var cs []*pms.ProductCategory
	// base category id is 0
	global.Db.Model(&pms.ProductCategory{}).Preload("Children").Where("parent_id  = 0").Find(&cs)
	common.Success(ctx, cs)

}

// Create
//
//	@Summary	create product category
//	@Tags		PmsProductCategoryController
//	@Accept		json
//	@Param		category	body		dto.CategoryDto	true	"add product category"
//	@Success	200			{object}	common.Result{data=pms.ProductCategory}
//	@Router		/productCategory/create [post]
func (p *productCategoryController) Create(ctx *gin.Context) {
	var cDto dto.CategoryDto
	if err := ctx.ShouldBindJSON(&cDto); err != nil {
		common.Error(ctx, "params error :"+err.Error())
		return
	}
	var c pms.ProductCategory
	err := copier.Copy(&c, &cDto)
	if err != nil {
		common.Error(ctx, "server error")
		return
	}

	if err := global.Db.Create(&c).Error; err != nil {

		common.Error(ctx, "create error: "+err.Error())
		return

	}
	common.Success(ctx, c)

}

// Page
//
//	@Summary	product category page view
//	@Tags		PmsProductCategoryController
//	@Param		parentId	path		int	true	"parent id"
//	@Success	200			{object}	common.PageResult{data=[]pms.ProductCategory}
//	@Router		/productCategory/list/{parentId} [get]
func (p *productCategoryController) Page(ctx *gin.Context) {
	parentId := ctx.Param("parentId")
	var cs []*pms.ProductCategory
	tx := global.Db.Model(&pms.ProductCategory{}).
		Where("parent_id = ?", parentId)
	var count int64
	tx.Count(&count)
	tx = tx.
		Scopes(common.PreparePage(ctx)).
		Find(&cs)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Page(ctx, cs, count)
}

// GetById
//
//	@Summary	get product category by id
//	@Tags		PmsProductCategoryController
//	@Param		id	path		int	true	"category id"
//	@Success	200	{object}	common.Result{data=pms.ProductCategory}
//	@Router		/productCategory/{id} [get]
func (p *productCategoryController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	var c pms.ProductCategory
	tx := global.Db.Model(&c).First(&c, id)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Success(ctx, c)
}

// DeleteById
//
//	@Summary	remove category by id
//	@Tags		PmsProductCategoryController
//	@Param		id	path		int	true	"wanted remove category id"
//	@Success	200	{object}	common.Result{data=int}
//	@Router		/productCategory/delete/{id} [post]
func (p *productCategoryController) DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")
	tx := global.Db.Delete(&pms.ProductCategory{}, id)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Success(ctx, tx.RowsAffected)
}

// UpdateById
//
//	@Summary	update category by id
//	@Tags		PmsProductCategoryController
//	@Param		id	path		int	true	"wanted update category's id"
//	@Success	200	{object}	common.Result{data=int}
//	@Router		/productCategory/update/{id} [post]
func (p *productCategoryController) UpdateById(ctx *gin.Context) {
	var cDto dto.CategoryDto
	var c pms.ProductCategory
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		common.Error(ctx, "id must be integer")
		return
	}
	if err := ctx.ShouldBindJSON(&cDto); err != nil {
		common.Error(ctx, "params error:"+err.Error())
		return
	}
	err = copier.Copy(&c, &cDto)
	if err != nil {
		common.Error(ctx, "server error")
		return
	}
	c.ID = uint(id)

	err = global.Db.Debug().Transaction(func(tx *gorm.DB) error {
		tx.Updates(&c)

		err2 := tx.Model(&c).Association("ProductAttributes").Replace(c.ProductAttributes)
		return err2
	})
	//tx := global.Db.Debug().Save(&c)

	//if tx.Error != nil || tx.RowsAffected == 0 {
	//	common.Error(ctx, "update error: "+err.Error())
	//	return
	//}
	common.Success(ctx, "ok")

}

// UpdateNavStatus
//
//	@Summary	update category's nav status when id in ids
//	@Tags		PmsProductCategoryController
//	@Param		navStatus	formData	int		true	"nav status"	Enums(0,1)
//	@Param		ids			formData	[]int	true	"id list"
//	@Success	200			{object}	common.Result{data=int}
//	@Router		/productCategory/update/navStatus  [post]
func (p *productCategoryController) UpdateNavStatus(ctx *gin.Context) {
	var params dto.UpdateCategoryNavStatusDto
	if err := ctx.ShouldBind(&params); err != nil {
		common.Error(ctx, "params error: "+err.Error())
		return
	}
	tx := global.Db.Model(&pms.ProductCategory{}).Where("id in ?", params.Ids).Update("nav_status", params.NavStatus)
	if tx.Error != nil || tx.RowsAffected == 0 {
		common.Error(ctx, "update error")
		return
	}
	common.Success(ctx, tx.RowsAffected)
}

// UpdateShowStatus
//
//	@Summary	update category's show status when id in ids
//	@Tags		PmsProductCategoryController
//	@Param		showStatus	formData	int		true	"nav status"	Enums(0,1)
//	@Param		ids			formData	[]int	true	"id list"
//	@Success	200			{object}	common.Result{data=int}
//	@Router		/productCategory/update/showStatus  [post]
func (p *productCategoryController) UpdateShowStatus(ctx *gin.Context) {
	var params dto.UpdateCategoryShowStatusDto
	if err := ctx.ShouldBind(&params); err != nil {
		common.Error(ctx, "params error: "+err.Error())
		return
	}
	tx := global.Db.Model(&pms.ProductCategory{}).Where("id in ?", params.Ids).Update("show_status", params.ShowStatus)
	if tx.Error != nil || tx.RowsAffected == 0 {
		common.Error(ctx, "update error")
		return
	}
	common.Success(ctx, tx.RowsAffected)
}
