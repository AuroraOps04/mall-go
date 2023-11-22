package pms

import (
	"github.com/gin-gonic/gin"
)

func Route(engine *gin.RouterGroup) {
	// brand manager route
	brandGroup := engine.Group("/brand")
	{
		brandGroup.GET("/:id", BrandController.GetById)
		brandGroup.GET("/list", BrandController.Page)
		brandGroup.GET("/listAll", BrandController.ListAll)
		brandGroup.POST("/create", BrandController.Create)
		brandGroup.GET("/delete/:id", BrandController.DeleteById)
		brandGroup.POST("/delete/batch", BrandController.DeleteBatch)
		brandGroup.POST("/update/factoryStatus", BrandController.UpdateFactoryStatusBatch)
		brandGroup.POST("/update/showStatus", BrandController.UpdateShowStatusBatch)
		brandGroup.POST("/update/:id", BrandController.Update)

	}
	// attribute manager route
	attributeGroup := engine.Group("/productAttribute")
	{
		attributeGroup.GET("/:id", ProductAttributeController.GetById)
		attributeGroup.GET("/attrInfo/:productCategoryId", ProductAttributeController.GetInfoByProductCategoryId)
		attributeGroup.POST("/create", ProductAttributeController.Create)
		attributeGroup.POST("/delete", ProductAttributeController.DeleteBatch)
		attributeGroup.GET("/list/:cId", ProductAttributeController.PageByProductAttributeCategoryId)
		attributeGroup.POST("/update/:id", ProductAttributeController.UpdateById)

		// attribute category manager route
		categoryGroup := attributeGroup.Group("/category")
		{
			categoryGroup.GET("/:id", AttributeCategoryController.GetById)
			categoryGroup.POST("/create", AttributeCategoryController.Create)
			categoryGroup.GET("/delete/:id", AttributeCategoryController.DeleteById)
			categoryGroup.GET("/list", AttributeCategoryController.Page)
			categoryGroup.GET("/list/withAttr", AttributeCategoryController.ListAllWithAttr)
			categoryGroup.POST("/update/:id", AttributeCategoryController.UpdateById)

		}

	}

	// product category manager route
	categoryGroup := engine.Group("/productCategory")
	{
		categoryGroup.GET("/:id", ProductCategoryController.GetById)
		categoryGroup.POST("/create", ProductCategoryController.Create)
		categoryGroup.POST("/delete/:id", ProductCategoryController.DeleteById)
		categoryGroup.GET("/list/:parentId", ProductCategoryController.Page)
		categoryGroup.GET("/list/withChildren", ProductCategoryController.ListAllWithChildren)
		categoryGroup.POST("/update/:id", ProductCategoryController.UpdateById)
		categoryGroup.POST("/update/navStatus", ProductCategoryController.UpdateNavStatus)
		categoryGroup.POST("/update/showStatus", ProductCategoryController.UpdateShowStatus)
	}

	productGroup := engine.Group("/product")
	{
		productGroup.GET("/list", ProductController.Page)
		productGroup.POST("/create", ProductController.Create)
		productGroup.POST("/update/newStatus", ProductController.UpdateNewStatus)
		productGroup.POST("/update/publishStatus", ProductController.UpdatePublishStatus)
		productGroup.POST("/update/recommendStatus", ProductController.UpdateRecommendStatus)
		productGroup.POST("/update/deleteStatus", ProductController.UpdateDeleteStatus)
		productGroup.POST("/update/verifyStatus", ProductController.UpdateVerifyStatus)
		productGroup.GET("/updateInfo/:id", ProductController.GetInfoForUpdate)
		productGroup.GET("/simpleList", ProductController.GetSimpleList)
		productGroup.POST("/update/:id", ProductController.UpdateById)

	}
}
