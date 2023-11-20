package pms

import "github.com/AuroraOps04/mall-go/model/base/field"

type PageProductDto struct {
	BrandId           *uint   `json:"brandId,omitempty"           form:"brandId"`
	Keyword           *string `json:"keyword,omitempty"           form:"keyword"`
	ProductCategoryId *uint   `json:"productCategoryId,omitempty" form:"productCategoryId"`
	ProductSN         *string `json:"productSN,omitempty"         form:"productSN"`
	PublishStatus     *int    `json:"publishStatus,omitempty"     form:"publishStatus"`
	VerifyStatus      *int    `json:"verifyStatus,omitempty"      form:"verifyStatus"`
}

type UpdateProductNewStatusDto struct {
	IDs       string                 `json:"ids"       form:"ids"       binding:"required"`
	NewStatus *field.LikeNumberInt64 `json:"newStatus" form:"newStatus" binding:"required"`
}

type UpdateProductPublishStatusDto struct {
	IDs           string                 `json:"ids"           form:"ids"           binding:"required"`
	PublishStatus *field.LikeNumberInt64 `json:"publishStatus" form:"publishStatus" binding:"required"`
}

type UpdateProductRecommendStatusDto struct {
	IDs             string                 `json:"ids"             form:"ids"             binding:"required"`
	RecommendStatus *field.LikeNumberInt64 `json:"recommendStatus" form:"recommendStatus" binding:"required"`
}

type UpdateProductDeleteStatusDto struct {
	IDs          string                 `json:"ids"          form:"ids"          binding:"required"`
	DeleteStatus *field.LikeNumberInt64 `json:"deleteStatus" form:"deleteStatus" binding:"required"`
}

type UpdateProductVerifyStatusDto struct {
	IDs          string                 `json:"ids"          form:"ids"          binding:"required"`
	VerifyStatus *field.LikeNumberInt64 `json:"verifyStatus" form:"verifyStatus" binding:"required"`
}
