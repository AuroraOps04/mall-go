package pms

import (
	"github.com/AuroraOps04/mall-go/model/base"
	"github.com/AuroraOps04/mall-go/model/base/field"
)

type ProductCategory struct {
	base.BaseModel
	ParentID          uint                   `json:"parentId"`
	Parent            *ProductCategory       `json:"parent"            gorm:""`
	Children          []*ProductCategory     `json:"children"          gorm:"foreignKey:ParentID"`
	Name              string                 `json:"name"              gorm:"size:64"`
	Level             int8                   `json:"level"             gorm:"size:1"`
	ProductCount      *field.LikeNumberInt64 `json:"productCount"`
	ProductUnit       string                 `json:"productUnit"       gorm:"size:64"`
	NavStatus         int8                   `json:"navStatus"         gorm:"size:1"`
	ShowStatus        int8                   `json:"showStatus"        gorm:"size:1"`
	Sort              *field.LikeNumberInt64 `json:"sort"`
	Icon              string                 `json:"icon"              gorm:"size:255"`
	Keywords          string                 `json:"keywords"          gorm:"size:255"`
	Description       string                 `json:"description"       gorm:"type:text"`
	ProductAttributes []*ProductAttribute    `json:"productAttributes" gorm:"many2many:pms_product_category_attribute_relation"`
}

func (p ProductCategory) TableName() string {
	return "pms_product_category"
}

type Product struct {
	base.BaseModel
	BrandID                    uint                     `json:"brandId"`
	Brand                      *Brand                   `json:"brand"`
	ProductCategoryID          uint                     `json:"productCategoryId"`
	CateParentID               uint                     `json:"cateParentId"`
	ProductCategory            ProductCategory          `json:"productCategory"`
	FeightTemplateID           uint                     `json:"feightTemplateId"`
	FeightTemplate             FeightTemplate           `json:"feightTemplate"`
	ProductAttributeCategoryID uint                     `json:"productAttributeCategoryId"`
	ProductAttributeCategory   ProductAttributeCategory `json:"productAttributeCategory"`
	Name                       string                   `json:"name"                       gorm:"size:64;index:index_name"`
	Pic                        string                   `json:"pic"                        gorm:"size:255"`
	ProductSn                  string                   `json:"productSn"                  gorm:"size:64;index:index_product_sn"`
	DeleteStatus               uint8                    `json:"deleteStatus"               gorm:"size:1"`
	PublishStatus              uint8                    `json:"publishStatus"              gorm:"size:1"`
	NewStatus                  uint8                    `json:"newStatus"                  gorm:"size:1"`
	RecommandStatus            uint8                    `json:"recommandStatus"            gorm:"size:1"`
	VerifyStatus               uint8                    `json:"verifyStatus"               gorm:"size:1"`
	Sort                       *field.LikeNumberInt64   `json:"sort"`
	Sale                       *field.LikeNumberInt64   `json:"sale"`
	Price                      *field.LikeNumberFloat64 `json:"price"                      gorm:"type:decimal(10,2)"`
	PromotionPrice             *string                  `json:"promotionPrice"             gorm:"type:decimal(10,2)"`
	GiftGrowth                 *field.LikeNumberInt64   `json:"giftGrowth"`
	GiftPoint                  *field.LikeNumberInt64   `json:"giftPoint"`
	UsePointLimit              *field.LikeNumberInt64   `json:"usePointLimit"`
	SubTitle                   string                   `json:"subTitle"                   gorm:"size:255"`
	Description                string                   `json:"description"                gorm:"type:text"`
	OriginalPrice              *field.LikeNumberFloat64 `json:"originalPrice"              gorm:"type:decimal(10,2)"`
	Stock                      *field.LikeNumberInt64   `json:"stock"`
	LowStock                   *field.LikeNumberInt64   `json:"lowStock"`
	Unit                       string                   `json:"unit"                       gorm:"size:16"`
	Weight                     *field.LikeNumberFloat64 `json:"weight"                     gorm:"type:decimal(10,2)"`
	PreviewStatus              uint8                    `json:"previewStatus"              gorm:"size:1"`
	ServiceIds                 string                   `json:"serviceIds"                 gorm:"size:64"`
	Keywords                   string                   `json:"keywords"                   gorm:"size:255"`
	Note                       string                   `json:"note"                       gorm:"size:255"`
	AlbumPics                  string                   `json:"albumPics"                  gorm:"size:255"`
	DetailTitle                string                   `json:"detailTitle"                gorm:"size:255"`
	DetailDesc                 string                   `json:"detailDesc"                 gorm:"type:text"`
	DetailHtml                 string                   `json:"detailHtml"                 gorm:"type:text"`
	DetailMobileHtml           string                   `json:"detailMobileHtml"           gorm:"type:text"`
	PromotionStartTime         *field.DateTime          `json:"promotionStartTime"                                               time_format:"2006-01-02" time_utc:"8"`
	PromotionEndTime           *field.DateTime          `json:"promotionEndTime"`
	PromotionPerLimit          *field.LikeNumberInt64   `json:"promotionPerLimit"`
	PromotionType              uint8                    `json:"promotionType"              gorm:"size:1"`
	ProductCategoryName        string                   `json:"productCategoryName"        gorm:"size:255"`
	BrandName                  string                   `json:"brandName"                  gorm:"size:255"`
	ProductAttributeValueList  []*ProductAttributeValue `json:"productAttributeValueList"`
	MemberPriceList            []*MemberPrice           `json:"memberPriceList"`
	ProductFullReductionList   []*ProductFullReduction  `json:"productFullReductionList"`
	ProductLadderList          []*ProductLadder         `json:"productLadderList"`
	SkuStockList               []*SkuStock              `json:"skuStockList"`
}

func (p Product) TableName() string {
	return "pms_product"
}

type ProductLadder struct {
	base.BaseModel
	ProductID uint                     `json:"productId"`
	Product   *Product                 `json:"product"`
	Count     *field.LikeNumberInt64   `json:"count"`
	Discount  *field.LikeNumberFloat64 `json:"discount"  gorm:"type:decimal(10,2)"`
	Price     *field.LikeNumberFloat64 `json:"price"     gorm:"type:decimal(10,2)"`
}

func (p ProductLadder) TableName() string {
	return "pms_product_ladder"
}

type ProductFullReduction struct {
	base.BaseModel
	ProductID   uint                     `json:"productId"`
	Product     *Product                 `json:"product"`
	FullPrice   *field.LikeNumberFloat64 `json:"fullPrice"`
	ReducePrice *field.LikeNumberFloat64 `json:"reducePrice"`
}

func (p ProductFullReduction) TableName() string {
	return "pms_product_full_reduction"
}

type ProductOperateLog struct {
	ProductID    uint                     `json:"productId"`
	Product      *Product                 `json:"product"`
	PriceOld     *field.LikeNumberFloat64 `json:"priceOld"      gorm:"type:decimal(10,2)"`
	PriceNew     *field.LikeNumberFloat64 `json:"priceNew"      gorm:"type:decimal(10,2)"`
	SalePriceOld *field.LikeNumberFloat64 `json:"salePriceOld"  gorm:"type:decimal(10,2)"`
	SalePriceNew *field.LikeNumberFloat64 `json:"salePriceNew"  gorm:"type:decimal(10,2)"`
	GiftPointOld *field.LikeNumberInt64   `json:"gift_pint_old"`
	GiftPointNew *field.LikeNumberInt64   `json:"giftPointNew"`
	UsePointOld  *field.LikeNumberInt64   `json:"usePointOld"`
	UsePointNew  *field.LikeNumberInt64   `json:"usePointNew"`
	OperateMan   string                   `json:"operateMan"    gorm:"size:64"`
}

type ProductVertifyRecord struct {
	ProductID  uint     `json:"productId"  gorm:""`
	Product    *Product `json:"product"`
	VertifyMan string   `json:"vertifyMan" gorm:"size:64"`
	Status     uint8    `json:"status"     gorm:"size:1"`
	Detail     string   `json:"detail"     gorm:"size:255"`
}
