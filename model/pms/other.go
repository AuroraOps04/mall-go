package pms

import (
	"github.com/AuroraOps04/mall-go/model/base"
)

type MemberPrice struct {
	base.BaseModel
	ProductID       uint    `json:"product_id"`
	MemberLevelID   uint    `json:"member_level_id"`
	MemberPrice     float64 `json:"member_price" gorm:"type:decimal(10,2)"`
	MemberLevelName string  `json:"member_level_name" gorm:"size:100"`
}

func (m MemberPrice) TableName() string {
	return "pms_member_price"
}

type FeightTemplate struct {
	base.BaseModel
	Name           string  `json:"name" gorm:"size:64"`
	ChargeType     int8    `json:"charge_type" gorm:"size:1"`
	FirstWeight    float64 `json:"first_weight" gorm:"type:decimal(10,2)"`
	FirstFee       float64 `json:"first_fee" gorm:"type:decimal(10,2)"`
	ContinueWeight float64 `json:"continue_weight" gorm:"type:decimal(10,2)"`
	ContinueFee    float64 `json:"continue_fee" gorm:"type:decimal(10,2)"`
	Dest           string  `json:"dest" gorm:"size:255"`
}

func (f FeightTemplate) TableName() string {
	return "pms_feight_template"
}

type SkuStock struct {
	base.BaseModel
	ProductID      uint    `json:"product_id"`
	SkuCode        string  `json:"sku_code" gorm:"size:64"`
	Price          float64 `json:"price" gorm:"type:decimal(10,2)"`
	Stock          uint    `json:"stock"`
	LowStock       int     `json:"low_stock"`
	Sp1            string  `json:"sp1" gorm:"size:64;column:sp1"`
	Sp2            string  `json:"sp2" gorm:"size:64;column:sp2"`
	Sp3            string  `json:"sp3" gorm:"size:64;column:sp3"`
	Pic            string  `json:"pic" gorm:"size:255"`
	Sale           int     `json:"sale"`
	PromotionPrice float64 `json:"promotion_price" gorm:"type:decimal(10,2)"`
	LockStock      int     `json:"lock_stock"`
}

func (s SkuStock) TableName() string {
	return "pms_sku_stock"
}
