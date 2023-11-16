package pms

import (
	"github.com/AuroraOps04/mall-go/common"
	dto "github.com/AuroraOps04/mall-go/dto/pms"
	"github.com/AuroraOps04/mall-go/global"
	"github.com/AuroraOps04/mall-go/model/pms"
	"github.com/gin-gonic/gin"
)

type productAttributeController struct {
}

var ProductAttributeController = &productAttributeController{}

// GetById
//
//	@Summary	get attribute by id
//	@Tags		PmsProductAttributeController
//	@Produce	json
//	@Param		id	path		int	true	"attribute id"
//	@Success	200	{object}	common.Result{data=pms.ProductAttribute}
//	@Router		/productAttribute/{id} [get]
func (p *productAttributeController) GetById(ctx *gin.Context) {
	var b pms.ProductAttribute
	id := ctx.Param("id")
	if err := global.Db.Model(&b).Where("id =?", id).First(&b).Error; err != nil {
		common.Error(ctx, "not found: "+err.Error())
		return
	}
	common.Success(ctx, b)

}

// GetInfoByProductCategoryId
//
//	@Summary	GetInfoByProductCategoryId
//	@Param		productCategoryId	path	int	true	"product category id"
//	@Tags		PmsProductAttributeController
//	@Success	200	{object}	common.Result{data=[]dto.AttrInfoDto}
//	@Router		/productAttribute/attrInfo/{productCategoryId} [get]
func (p *productAttributeController) GetInfoByProductCategoryId(ctx *gin.Context) {
	productCategoryId := ctx.Param("productCategoryId")
	// productCategory and productAttribute is many2many
	var dtoList []*dto.AttrInfoDto
	tx := global.Db.Raw(`select ppa.id as 'attribute_id', ppa.product_attribute_category_id as 'attribute_category_id'
				from pms_product_attribute ppa
        		 left join pms_product_category_attribute_relation ppcar on ppcar.product_attribute_id = ppa.id
				where ppcar.product_category_id = ?
				`, productCategoryId).
		Scan(&dtoList)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Success(ctx, dtoList)

}

// Create
//
//	@Summary	create attribute
//	@Param		attribute	body	pms.ProductAttribute	true	"add attribute"
//	@Tags		PmsProductAttributeController
//	@Success	200	{object}	common.Result{data=pms.ProductAttribute}
//	@Router		/productAttribute/create [post]
func (p *productAttributeController) Create(ctx *gin.Context) {
	var attr pms.ProductAttribute
	if err := ctx.ShouldBindJSON(&attr); err != nil {
		common.Error(ctx, "params error: "+err.Error())
		return
	}
	tx := global.Db.Create(&attr)
	if tx.Error != nil {
		common.Error(ctx, "create error: "+tx.Error.Error())
		return
	}
	common.Success(ctx, attr)
}

// DeleteBatch
//
//	@Summary	Delete Batch productAttribute
//	@Tags		PmsProductAttributeController,uncompleted
//	@Router		/productAttribute/delete [post]
func (p *productAttributeController) DeleteBatch(ctx *gin.Context) {

}

// PageByProductAttributeCategoryId
//
//	@Summary	list by page
//	@Tags		PmsProductAttributeController
//	@Param		cId		path		int	true	"product attribute category id"
//	@Param		type	query		int	true	"type"
//	@Success	200		{object}	common.PageResult{data=[]pms.ProductAttribute}
//	@Router		/productAttribute/list/{cId} [get]
func (p *productAttributeController) PageByProductAttributeCategoryId(ctx *gin.Context) {
	// product attribute category id
	cId := ctx.Param("cId")
	// 0 is attribute 1 is parameter
	aType := ctx.Query("type")

	var as []*pms.ProductAttribute

	tx := global.Db.Model(&pms.ProductAttribute{}).
		Where("type = ? and product_attribute_category_id = ?", aType, cId)

	var count int64
	tx.Count(&count)
	tx = tx.Scopes(common.PreparePage(ctx)).Find(&as)

	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Page(ctx, as, count)

}

// UpdateById
//
//	@Summary	update attribute by id
//	@Tags		PmsProductAttributeController
//	@Param		id		path		int						true	"attribute id"
//	@Param		attr	body		pms.ProductAttribute	true	"update attribute "
//	@Success	200		{object}	common.Result{data=int}
//	@Router		/productAttribute/{id} [post]
func (p *productAttributeController) UpdateById(ctx *gin.Context) {
	id := ctx.Param("id")
	var b pms.ProductAttribute
	if err := ctx.ShouldBindJSON(&b); err != nil {
		common.Error(ctx, "params error: "+err.Error())
		return
	}

	tx := global.Db.Debug().Where("id = ?", id).Save(&b)
	if tx.Error != nil {
		common.Error(ctx, "update error : "+tx.Error.Error())
		return
	}
	common.Success(ctx, tx.RowsAffected)

}
