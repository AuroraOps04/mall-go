package pms

import (
	"fmt"

	"github.com/AuroraOps04/mall-go/common"
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
	tx := global.Db.Model(&pms.Product{})
	if param.BrandId != nil {
		tx = tx.Where("brand_id = ?", param.BrandId)
	}
	if param.Keyword != nil {
		tx = tx.Where("keyword like  ?", "%"+*param.Keyword+"%")
	}
	if param.ProductCategoryId != nil {
		tx = tx.Where("product_category_id = ?", param.ProductCategoryId)
	}
	if param.ProductSN != nil {
		tx = tx.Where("product_sn = ?", param.ProductSN)
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
//	@Summary	Create prodcut api
//	@Tags		PmsProductController
//	@Param		data body pms.Product true "add product"
//	@Router		/product/create [post]
func (p *productController) Create(ctx *gin.Context) {
	var dto pms.Product
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		fmt.Printf("%t\n", err)
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
