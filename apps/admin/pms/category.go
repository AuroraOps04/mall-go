package pms

import (
	"github.com/AuroraOps04/mall-go/common"
	"github.com/AuroraOps04/mall-go/global"
	"github.com/AuroraOps04/mall-go/model/base"
	"github.com/AuroraOps04/mall-go/model/pms"
	"github.com/gin-gonic/gin"
	"strconv"
)

type productCategoryController struct {
}

var ProductCategoryController = &productCategoryController{}

func (p *productCategoryController) base(ctx *gin.Context) {

}

func (p *productCategoryController) ListAllWithChildren(ctx *gin.Context) {
	var cs []*pms.ProductCategory
	// base category id is 0
	global.Db.Model(&pms.ProductCategory{}).Preload("Children").Where("parent_id  = 0").Find(&cs)
	common.Success(ctx, cs)

}

func (p *productCategoryController) Create(ctx *gin.Context) {
	var c pms.ProductCategory
	if err := ctx.ShouldBindJSON(&c); err != nil {
		common.Error(ctx, "params error :"+err.Error())
		return
	}

	attributeIds := ctx.PostFormArray("productAttributeIdList")
	var attributes []*pms.ProductAttribute
	for _, id := range attributeIds {
		parseUint, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			continue
		}
		attributes = append(attributes, &pms.ProductAttribute{
			BaseModel: base.BaseModel{
				ID: uint(parseUint),
			},
		})
	}

	c.ProductAttributes = attributes
	if err := global.Db.Create(&c).Error; err != nil {

		common.Error(ctx, "create error: "+err.Error())
		return

	}
	common.Success(ctx, c)

}

func (p *productCategoryController) Page(ctx *gin.Context) {
	parentId := ctx.Param("parentId")
	var cs []*pms.ProductCategory
	tx := global.Db.Model(&pms.ProductCategory{}).
		Scopes(common.PreparePage(ctx)).
		Where("parent_id = ?", parentId)
	var count int64
	tx.Count(&count)
	tx.Find(&cs)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Page(ctx, cs, count)
}
