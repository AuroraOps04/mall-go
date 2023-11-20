package pms

import (
	"fmt"

	"github.com/AuroraOps04/mall-go/common"
	"github.com/AuroraOps04/mall-go/common/util"
	dto "github.com/AuroraOps04/mall-go/dto/pms"
	"github.com/AuroraOps04/mall-go/global"
	"github.com/AuroraOps04/mall-go/model/pms"
	"github.com/gin-gonic/gin"
)

type productController struct {
}

var ProductController = &productController{}

// Page
//
//	@Summary	Page
//	@Param		brandId				query	int		false	"brand id"
//	@Param		keyword				query	string	false	"keyword"
//	@Param		pageNum				query	int		false	"page num"	default(1)
//	@Param		pageSize			query	int		false	"page size"	default(1)
//	@Param		productCategoryId	query	int		false	"product category id"
//	@Param		productSN			query	string	false	"product sn"
//	@Param		publishStatus		query	int		false	"publish status"
//	@Param		verifyStatus		query	int		false	"verify status"
//	@Tags		PmsProductController
//	@Success	200	{object}	common.PageResult{}
//	@Router		/product/list [get]
func (p *productController) Page(ctx *gin.Context) {
	var param dto.PageProductDto
	if err := ctx.ShouldBindQuery(&param); err != nil {
		common.Error(ctx, "params err: "+err.Error())
		return
	}
	fmt.Println("page product params: ", param)
	tx := global.Db.Model(&pms.Product{})
	if param.BrandId != nil {
		tx = tx.Where("brand_id = ?", param.BrandId)
	}
	if param.Keyword != nil {
		tx = tx.Where("keywords like  ?", "%"+*param.Keyword+"%")
	}
	if param.ProductCategoryId != nil {
		tx = tx.Where("product_category_id = ?", param.ProductCategoryId)
	}
	if param.ProductSN != nil {
		tx = tx.Where("product_sn like ?", "%"+*param.ProductSN+"%")
	}
	if param.PublishStatus != nil {
		tx = tx.Where("publish_status = ?", param.PublishStatus)
	}
	if param.VerifyStatus != nil {
		tx = tx.Where("verify_status = ?", param.VerifyStatus)
	}

	var data []*pms.Product
	var count int64
	tx.Count(&count)
	tx = tx.Scopes(common.PreparePage(ctx)).Order("sort desc").Find(&data)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Page(ctx, data, count)

}

// Create
//
//	@Summary	Create prodcut api
//	@Tags		PmsProductController
//	@Param		data	body	pms.Product	true	"add product"
//	@Success	200		object	common.Result{data=int}
//	@Router		/product/create [post]
func (p *productController) Create(ctx *gin.Context) {
	var dto pms.Product
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		fmt.Println(err)
		common.Error(ctx, err.Error())
		return
	}
	// TODO: process realtions
	tx := global.Db.Create(&dto)
	if tx.Error != nil || tx.RowsAffected == 0 {
		common.Error(ctx, "create error")
		return
	}
	common.Success(ctx, tx.RowsAffected)

}

// UpdateNewStatus
//
//	@Summary	updateProductNewStatus
//	@Tags		PmsProductController
//	@Param		ids			query	string	true	"id list"
//	@Param		newStatus	query	string	true	"newStatus"	Enums(0,1)
//	@Success	200			object	common.Result{data=int}
//	@Router		/product/update/newStatus [post]
func (p *productController) UpdateNewStatus(ctx *gin.Context) {
	var param dto.UpdateProductNewStatusDto

	if err := ctx.ShouldBindQuery(&param); err != nil {
		common.Error(ctx, err.Error())
		return
	}
	ids := util.Str2IntSlice(&param.IDs, ",")
	tx := global.Db.Model(&pms.Product{}).
		Where("id in ?", ids).
		Update("new_status = ?", param.NewStatus)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Success(ctx, tx.RowsAffected)
}

// UpdatePublishStatus
//
//	@Summary	updateProductPublishStatus
//	@Tags		PmsProductController
//	@Param		ids				query	string	true	"id list"
//	@Param		publishStatus	query	int		true	"publishStatus"	Enums(0,1)
//	@Success	200				object	common.Result{data=int}
//	@Router		/product/update/publishStatus [post]
func (p *productController) UpdatePublishStatus(ctx *gin.Context) {

	var param dto.UpdateProductPublishStatusDto

	if err := ctx.ShouldBindQuery(&param); err != nil {
		common.Error(ctx, err.Error())
		return
	}
	ids := util.Str2IntSlice(&param.IDs, ",")
	tx := global.Db.Model(&pms.Product{}).
		Where("id in ?", ids).
		Update("publish_status = ?", param.PublishStatus)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Success(ctx, tx.RowsAffected)
}

// UpdateRecommendStatus
//
//	@Summary	updateProductRecommendStatus
//	@Tags		PmsProductController
//	@Param		ids				query	string	true	"id list"
//	@Param		recommendStatus	query	int		true	"recommendStatus"	Enums(0,1)
//	@Success	200				object	common.Result{data=int}
//	@Router		/product/update/recommendStatus [post]
func (p *productController) UpdateRecommendStatus(ctx *gin.Context) {
	var param dto.UpdateProductRecommendStatusDto

	if err := ctx.ShouldBindQuery(&param); err != nil {
		common.Error(ctx, err.Error())
		return
	}
	ids := util.Str2IntSlice(&param.IDs, ",")
	tx := global.Db.Model(&pms.Product{}).
		Where("id in ?", ids).
		Update("recommend_status = ?", param.RecommendStatus)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Success(ctx, tx.RowsAffected)
}

// UpdateDeleteStatus
//
//	@Summary	updateProductDeleteStatus
//	@Tags		PmsProductController
//	@Param		ids				query	string	true	"id list"
//	@Param		deleteStatus	query	int		true	"deleteStatus"	Enums(0,1)
//	@Success	200				object	common.Result{data=int}
//	@Router		/product/update/deleteStatus [post]
func (p *productController) UpdateDeleteStatus(ctx *gin.Context) {
	var param dto.UpdateProductDeleteStatusDto

	if err := ctx.ShouldBindQuery(&param); err != nil {
		common.Error(ctx, err.Error())
		return
	}
	ids := util.Str2IntSlice(&param.IDs, ",")
	tx := global.Db.Model(&pms.Product{}).
		Where("id in ?", ids).
		Update("delete_status = ?", param.DeleteStatus)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Success(ctx, tx.RowsAffected)
}

// UpdateVerifyStatus
//
//	@Summary	updateProductVerifyStatus
//	@Tags		PmsProductController
//	@Param		ids				query	string	true	"id list"
//	@Param		deleteStatus	query	int		true	"deleteStatus"	Enums(0,1)
//	@Success	200				object	common.Result{data=int}
//	@Router		/product/update/verifyStatus [post]
func (p *productController) UpdateVerifyStatus(ctx *gin.Context) {
	var param dto.UpdateProductVerifyStatusDto

	if err := ctx.ShouldBindQuery(&param); err != nil {
		common.Error(ctx, err.Error())
		return
	}
	ids := util.Str2IntSlice(&param.IDs, ",")
	tx := global.Db.Model(&pms.Product{}).
		Where("id in ?", ids).
		Update("verfiy_status = ?", param.VerifyStatus)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Success(ctx, tx.RowsAffected)
}

// GetInfoForUpdate
//
//	@Summary	getProductInfoForUpdate
//	@Tags		PmsProductController
//	@Success	200	object	common.Result
//	@Router		/product/updateInfo/{id} [get]
func (p *productController) GetInfoForUpdate(ctx *gin.Context) {

}
