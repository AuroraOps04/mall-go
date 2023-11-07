package pms

import (
	"github.com/AuroraOps04/mall-go/model/base"
)

type ProductAttributeCategory struct {
	base.BaseModel
	Name              string              `json:"name" gorm:"size:64"`
	AttributeCount    uint                `json:"attribute_count"`
	ParamCount        uint                `json:"param_count"`
	ProductAttributes []*ProductAttribute `json:"product_attributes" gorm:"foreignKey:ProductAttributeCategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (p ProductAttributeCategory) TableName() string {
	return "pms_product_attribute_category"
}

type ProductAttribute struct {
	base.BaseModel
	ProductAttributeCategoryID uint                      `json:"product_attribute_category_id"`
	ProductAttributeCategory   *ProductAttributeCategory `json:"product_attribute_category" gorm:"foreignKey:ProductAttributeCategoryID"`
	Name                       string                    `json:"name" gorm:"size:64"`
	SelectType                 int8                      `json:"select_type" gorm:"size:1"`
	InputType                  int8                      `json:"input_type" gorm:"size:1"`
	InputList                  string                    `json:"input_list" gorm:"size:255"`
	Sort                       int                       `json:"sort"`
	FilterType                 int8                      `json:"filter_type" gorm:"size:1"`
	SearchType                 int8                      `json:"search_type" gorm:"size:1"`
	RelatedStatus              int8                      `json:"related_status" gorm:"size:1"`
	HandAddStatus              int8                      `json:"hand_add_status" gorm:"size:1"`
	Type                       int8                      `json:"type" gorm:"size:1"`
	AttributeValues            []*ProductAttributeValue  `json:"attribute_values" gorm:"foreignKey:ProductAttributeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProductCategories          []*ProductCategory        `json:"product_categories" gorm:"many2many:pms_product_category_attribute_relation;"`
}

func (p ProductAttribute) TableName() string {
	return "pms_product_attribute"
}

type ProductAttributeValue struct {
	base.BaseModel
	ProductAttributeID uint              `json:"product_attribute_id"`
	ProductAttribute   *ProductAttribute `json:"product_attribute" gorm:"foreignKey:ProductAttributeID"`
	Value              string            `json:"value" gorm:"size:1000"`
	ProductId          uint              `json:"product_id"`
}

func (p ProductAttributeValue) TableName() string {
	return "pms_product_attribute_value"
}
