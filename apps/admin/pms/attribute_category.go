package pms

import (
	"github.com/AuroraOps04/mall-go/common"
	"github.com/AuroraOps04/mall-go/global"
	"github.com/AuroraOps04/mall-go/model/pms"
	"github.com/gin-gonic/gin"
)

type attributeCategoryController struct{}

var AttributeCategoryController = &attributeCategoryController{}

func (a *attributeCategoryController) Page(ctx *gin.Context) {
	var categories []*pms.ProductAttributeCategory
	tx := global.Db.Model(&pms.ProductAttributeCategory{}).Scopes(common.PreparePage(ctx))

	var count int64
	tx.Count(&count)
	tx.Find(&categories)
	common.Page(ctx, categories, count)
}

func (a *attributeCategoryController) Create(ctx *gin.Context) {
	name := ctx.PostForm("name")
	tx := global.Db.Create(&pms.ProductAttributeCategory{
		Name: name,
	})
	if tx.Error != nil {
		common.Success(ctx, tx.Error.Error())
		return
	}
	common.Success(ctx, name)
}

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

func (a *attributeCategoryController) DeleteById(ctx *gin.Context) {

	id := ctx.Param("id")
	tx := global.Db.Delete(&pms.ProductAttributeCategory{}, id)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}
	common.Success(ctx, tx.RowsAffected)
}

func (a *attributeCategoryController) ListAllWithAttr(ctx *gin.Context) {
	var list []*pms.ProductAttributeCategory
	global.Db.Debug().Model(&pms.ProductAttributeCategory{}).Preload("ProductAttributes").Find(&list)
	common.Success(ctx, list)
}

func (a *attributeCategoryController) UpdateById(ctx *gin.Context) {
	id := ctx.Param("id")
	var ac pms.ProductAttributeCategory

	if err := ctx.ShouldBindJSON(&ac); err != nil {
		common.Error(ctx, err.Error())
		return
	}

	tx := global.Db.Model(&ac).Where("id =?", id).Updates(ac)
	if tx.Error != nil {
		common.Error(ctx, tx.Error.Error())
		return
	}

	common.Success(ctx, ac)

}
