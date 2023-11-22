package pms

import (
	"fmt"
	"strconv"
	"time"

	"github.com/AuroraOps04/mall-go/common"
	"github.com/AuroraOps04/mall-go/common/util"
	dto "github.com/AuroraOps04/mall-go/dto/pms"
	"github.com/AuroraOps04/mall-go/global"
	"github.com/AuroraOps04/mall-go/model/base/field"
	"github.com/AuroraOps04/mall-go/model/pms"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type productController struct {
}

var ProductController = &productController{}

// setSkuStockCodeAndRelation
// SkuCode's format patten is
func setSkuStockCodeAndRelation(list []*pms.SkuStock, productId uint) {
	if len(list) == 0 {
		return
	}
	date := time.Now().Format("20060102")
	for i, v := range list {
		v.ProductID = productId
		if v.SkuCode == nil {
			s := fmt.Sprintf("%s%04d%03d", date, productId, i+1)
			v.SkuCode = &s

		}
	}
}

func getProductStock(list []*pms.SkuStock) *field.LikeNumberInt64 {
	var count uint = 0
	for _, v := range list {
		count += v.Stock
	}
	c := field.LikeNumberInt64(count)
	return &c
}

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
	tx := global.Db.Model(&pms.Product{}).
		Preload("ProductAttributeValueList").
		Preload("MemberPriceList").
		Preload("ProductFullReductionList").
		Preload("ProductLadderList").
		Preload("SkuStockList")

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
	err := global.Db.Transaction(func(tx *gorm.DB) error {
		skuStockList := dto.SkuStockList
		dto.SkuStockList = nil
		dto.Stock = getProductStock(skuStockList)
		if err := tx.Create(&dto).Error; err != nil {
			return err
		}
		setSkuStockCodeAndRelation(skuStockList, dto.ID)
		if len(skuStockList) == 0 {
			return nil
		}
		return tx.Where("id = ?", dto.ID).Association("skuStockList").Append(skuStockList)
	})

	if err != nil {
		common.Error(ctx, "create error: "+err.Error())
		return
	}
	common.Success(ctx, 1)

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
		Update("new_status", param.NewStatus)
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
		Update("publish_status", param.PublishStatus)
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
		Update("recommand_status", param.RecommendStatus)
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
		Update("delete_status", param.DeleteStatus)
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
		Update("verfiy_status", param.VerifyStatus)
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
//	@Param		id	path	int	true	"prodcut id"
//	@Success	200	object	common.Result{data=pms.Product}
//	@Router		/product/updateInfo/{id} [get]
func (p *productController) GetInfoForUpdate(ctx *gin.Context) {
	id := ctx.Param("id")
	var product pms.Product
	// FIXME: recursion call
	tx := global.Db.Debug().Model(&product).
		Preload(clause.Associations).
		Where("id = ?", id).
		First(&product)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	product.CateParentID = product.ProductCategory.ParentID
	fmt.Printf("%%#v: %#v\n%%v:%v\n", product, product)
	common.Success(ctx, product)

}

// GetSimpleList
//
//	@Tags		PmsProductController
//	@Summary	GetProductSimpleListNotContainsRelationData
//	@Param		keyword	query	string	false	"name or sn"
//	@Success	200		object	common.Result{data=[]pms.Product}
//	@Router		/product/simpleList [get]
func (p *productController) GetSimpleList(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	var list []*pms.Product
	keyword = "%" + keyword + "%"
	tx := global.Db.Model(&pms.Product{}).
		Where("name like ? or product_sn like", keyword, keyword).
		Order("sort desc").
		Find(&list)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Success(ctx, list)
}

// UpdateById
//
//	@Summary	PmsUpdateProductById
//	@Tags		PmsProductController
//	@Param		id		path	int			true	"product id"
//	@Param		prodcut	body	pms.Product	true	"want to update product"
//	@Success	200		object	common.Result{data=int}
//	@Router		/update/{id} [post]
func (p *productController) UpdateById(ctx *gin.Context) {
	id := ctx.Param("id")
	var product pms.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		common.Error(ctx, "params error: "+err.Error())
		return
	}
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		common.Error(ctx, "id must be int")
		return
	}
	product.ID = uint(idUint)
	err = global.Db.Transaction(func(tx *gorm.DB) error {
		// process relations
		// prepare skuStockList
		setSkuStockCodeAndRelation(product.SkuStockList, product.ID)
		// NOTE: wantch how to process relactions in gorm
		return tx.Debug().Where("id = ?", idUint).Save(&product).Error
	})
	if err != nil {
		common.Error(ctx, "update error: "+err.Error())
		return
	}
	common.Success(ctx, 1)
}
