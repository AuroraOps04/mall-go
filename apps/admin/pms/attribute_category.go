package pms

import (
	"github.com/AuroraOps04/mall-go/common"
	"github.com/AuroraOps04/mall-go/global"
	"github.com/AuroraOps04/mall-go/model/pms"
	"github.com/gin-gonic/gin"
)

type attributeCategoryController struct{}

var AttributeCategoryController = &attributeCategoryController{}

// Page
//
//	@Summary	list by page
//	@Tags		attribute category
//	@Success	200	{object}	common.PageResult{data=[]pms.ProductAttributeCategory}
//	@Router		/productAttribute/category/list [get]
func (a *attributeCategoryController) Page(ctx *gin.Context) {
	var categories []*pms.ProductAttributeCategory
	tx := global.Db.Model(&pms.ProductAttributeCategory{}).Scopes(common.PreparePage(ctx))

	var count int64
	tx.Count(&count)
	tx.Find(&categories)
	common.Page(ctx, categories, count)
}

// Create
//
//	@Summary	create attribute category
//	@Tags		attribute category
//	@Accept		json
//	@Param		name	formData	string	true	"category name"
//	@Success	200		{object}	common.Result{data=int}
//	@Router		/productAttribute/category/create [post]
func (a *attributeCategoryController) Create(ctx *gin.Context) {
	name := ctx.PostForm("name")
	tx := global.Db.Create(&pms.ProductAttributeCategory{
		Name: name,
	})
	if tx.Error != nil {
		common.Success(ctx, tx.Error.Error())
		return
	}
	common.Success(ctx, tx.RowsAffected)
}

// GetById
//
//	@Summary	get attribute category by id
//	@Tags		attribute category
//	@Param		id	path		int	true	"category id"
//	@Success	200	{object}	common.Result{data=pms.ProductAttributeCategory}
//	@Router		/productAttribute/category/{id} [get]
func (a *attributeCategoryController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	var ac pms.ProductAttributeCategory
	tx := global.Db.Model(ac).Where("id =?", id).First(&ac)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Success(ctx, ac)
}

// DeleteById
//
//	@Summary	Delete by  id
//	@Tags		attribute category
//	@Param		id	path		int	true	"category id"
//	@Success	200	{object}	common.Result{data=int}
//	@Router		/productAttribute/category/delete/{id} [get]
func (a *attributeCategoryController) DeleteById(ctx *gin.Context) {

	id := ctx.Param("id")
	tx := global.Db.Delete(&pms.ProductAttributeCategory{}, id)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Success(ctx, tx.RowsAffected)
}

// ListAllWithAttr
//
//	@Summary	select all and preload attribute list
//	@Tags		attribute category
//	@Success	200	{object}	common.Result{data=pms.ProductAttributeCategory}
//	@Router		/productAttribute/category/list/withAttr [get]
func (a *attributeCategoryController) ListAllWithAttr(ctx *gin.Context) {
	var list []*pms.ProductAttributeCategory
	global.Db.Model(&pms.ProductAttributeCategory{}).Preload("ProductAttributeList").Find(&list)
	common.Success(ctx, list)
}

// UpdateById
//
//	@Summary	update by  id
//	@Tags		attribute category
//	@Param		id			path		int								true	"category id"
//	@Param		category	body		pms.ProductAttributeCategory	true	"update category"
//	@Success	200			{object}	common.Result{data=int}
//	@Router		/productAttribute/category/update/{id} [post]
func (a *attributeCategoryController) UpdateById(ctx *gin.Context) {
	id := ctx.Param("id")
	var ac pms.ProductAttributeCategory

	if err := ctx.ShouldBindJSON(&ac); err != nil {
		common.Error(ctx, err.Error())
		return
	}

	tx := global.Db.Where("id =?", id).Updates(&ac)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}

	common.Success(ctx, tx.RowsAffected)

}
