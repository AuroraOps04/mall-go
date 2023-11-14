package pms

import (
	"github.com/AuroraOps04/mall-go/common/field"
	"github.com/AuroraOps04/mall-go/model/base"
)

type ProductAttributeCategory struct {
	base.BaseModel
	Name                 string              `json:"name" gorm:"size:64"`
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
	ProductAttributeCategory   *ProductAttributeCategory `json:"productAttributeCategory" gorm:"foreignKey:ProductAttributeCategoryID"`
	Name                       string                    `json:"name" gorm:"size:64"`
	SelectType                 int8                      `json:"selectType" gorm:"size:1"`
	InputType                  int8                      `json:"inputType" gorm:"size:1"`
	InputList                  string                    `json:"inputList" gorm:"size:255"`
	Sort                       field.LikeNumber          `json:"sort" gorm:"size:1"`
	FilterType                 int8                      `json:"filterType" gorm:"size:1"`
	SearchType                 int8                      `json:"searchType" gorm:"size:1"`
	RelatedStatus              int8                      `json:"relatedStatus" gorm:"size:1"`
	HandAddStatus              int8                      `json:"handAddStatus" gorm:"size:1"`
	Type                       int8                      `json:"type" gorm:"size:1"`
	AttributeValues            []*ProductAttributeValue  `json:"attributeValues" gorm:"foreignKey:ProductAttributeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ProductCategories          []*ProductCategory        `json:"productCategories" gorm:"many2many:pms_product_category_attribute_relation;"`
}

// TODO: 解决 JSON 从 数字 和字符串类型的数字 的字段反序列化

type Sort int8

func (s Sort) UnmarshalJSON(bytes []byte) error {
	return nil
}

func (s Sort) MarshalJSON() ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductAttribute) TableName() string {
	return "pms_product_attribute"
}

type ProductAttributeValue struct {
	base.BaseModel
	ProductAttributeID uint              `json:"productAttributeId"`
	ProductAttribute   *ProductAttribute `json:"productAttribute" gorm:"foreignKey:ProductAttributeID"`
	Value              string            `json:"value" gorm:"size:1000"`
	ProductId          uint              `json:"productId"`
}

func (p ProductAttributeValue) TableName() string {
	return "pms_product_attribute_value"
}
