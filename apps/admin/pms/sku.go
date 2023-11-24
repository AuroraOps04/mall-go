package pms

import (
	"strconv"

	"github.com/AuroraOps04/mall-go/common"
	"github.com/AuroraOps04/mall-go/global"
	"github.com/AuroraOps04/mall-go/model/pms"
	"github.com/gin-gonic/gin"
)

type skuStockController struct{}

var SkuStockController = &skuStockController{}

// GetByProductId
//
//	@Summary	get sku stock by prodcut id
//	@Tags		PmsSkuStockController
//	@Param		pid		path	int		true	"product id"
//	@Param		keyword	query	string	false	"sku code"
//	@Success	200		object	common.Result{data=[]pms.SkuStock}
//	@Router		/sku/{pid} [get]
func (p *skuStockController) GetByProductId(ctx *gin.Context) {
	pid := ctx.Param("pid")
	keyword := ctx.Query("keyword")
	tx := global.Db.Model(&pms.SkuStock{}).Where("product_id = ?", pid)
	if keyword != "" {
		tx = tx.Where("sku_code = ?", "%"+keyword+"%")
	}
	var skuList []pms.SkuStock
	if err := tx.Find(&skuList).Error; err != nil {
		common.Error(ctx, err.Error())
		return
	}
	common.Success(ctx, skuList)
}

// UpdateSkusById
//
//	@Summary	update product's skuList by product id
//	@Tags		PmsSkuStockController
//	@Param		pid		path	int				true	"prodcut id"
//	@Param		body	body	[]pms.SkuStock	true	"skuList"
//	@Success	200		object	common.Result{data=int}
//	@Router		/sku/update/{pid} [post]
func (p *skuStockController) UpdateSkusById(ctx *gin.Context) {
	pid := ctx.Param("pid")
	var skuList []*pms.SkuStock
	if err := ctx.ShouldBindJSON(&skuList); err != nil {
		common.Error(ctx, "params error: "+err.Error())
		return
	}
	id, err := strconv.ParseUint(pid, 10, 64)
	if err != nil {
		common.Error(ctx, "pid must be int")
		return
	}
	// this is correct
	product := pms.Product{}
	product.ID = uint(id)
	err = global.Db.Debug().Model(&product).
		// this is a mistake
		// Where("id = ?", pid).
		Association("SkuStockList").
		Replace(&skuList)
	if err != nil {
		common.Error(ctx, "update error: "+err.Error())
		return
	}

	common.Success(ctx, len(skuList))
}
