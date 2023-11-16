package pms

type PageProductDto struct {
	BrandId           *uint   `json:"brandId,omitempty"`
	Keyword           *string `json:"keyword,omitempty"`
	ProductCategoryId *uint   `json:"productCategoryId,omitempty"`
	ProductSN         *string `json:"productSN,omitempty"`
	PublishStatus     *int    `json:"publishStatus,omitempty"`
	VerifyStatus      *int    `json:"verifyStatus,omitempty"`
}
