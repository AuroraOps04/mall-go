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

	// product attribute manager route
	categoryGroup := engine.Group("/productCategory")
	{
		categoryGroup.POST("/create", ProductCategoryController.Create)
		categoryGroup.GET("/list/:parentId", ProductCategoryController.Page)
		categoryGroup.GET("/list/withChildren", ProductCategoryController.ListAllWithChildren)
	}

}
