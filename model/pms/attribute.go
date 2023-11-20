package pms

import (
	"github.com/AuroraOps04/mall-go/model/base"
	"github.com/AuroraOps04/mall-go/model/base/field"
)

type ProductAttributeCategory struct {
	base.BaseModel
	Name                 string              `json:"name"                 gorm:"size:64"`
	AttributeCount       uint                `json:"attributeCount"`
	ParamCount           uint                `json:"paramCount"`
	ProductAttributeList []*ProductAttribute `json:"productAttributeList" gorm:"foreignKey:ProductAttributeCategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (p ProductAttributeCategory) TableName() string {
	return "pms_product_attribute_category"
}

type ProductAttribute struct {
	base.BaseModel
	ProductAttributeCategoryID uint                      `json:"productAttributeCategoryId"`
	ProductAttributeCategory   *ProductAttributeCategory `json:"productAttributeCategory"   gorm:"foreignKey:ProductAttributeCategoryID"`
	Name                       string                    `json:"name"                       gorm:"size:64"`
	SelectType                 int8                      `json:"selectType"                 gorm:"size:1"`
	InputType                  int8                      `json:"inputType"                  gorm:"size:1"`
	InputList                  string                    `json:"inputList"                  gorm:"size:255"`
	Sort                       field.LikeNumberInt64          `json:"sort"                       gorm:"size:1"`
	FilterType                 int8                      `json:"filterType"                 gorm:"size:1"`
	SearchType                 int8                      `json:"searchType"                 gorm:"size:1"`
	RelatedStatus              int8                      `json:"relatedStatus"              gorm:"size:1"`
	HandAddStatus              int8                      `json:"handAddStatus"              gorm:"size:1"`
	Type                       int8                      `json:"type"                       gorm:"size:1"`
	AttributeValues            []*ProductAttributeValue  `json:"attributeValues"            gorm:"foreignKey:ProductAttributeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProductCategories          []*ProductCategory        `json:"productCategories"          gorm:"many2many:pms_product_category_attribute_relation;"`
}



type ProductAttributeValue struct {
	base.BaseModel
	ProductAttributeID uint              `json:"productAttributeId"`
	ProductAttribute   *ProductAttribute `json:"productAttribute"   gorm:"foreignKey:ProductAttributeID"`
	Value              string            `json:"value"              gorm:"size:1000"`
	ProductId          uint              `json:"productId"`
}

func (p ProductAttributeValue) TableName() string {
	return "pms_product_attribute_value"
}
