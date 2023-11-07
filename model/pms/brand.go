package pms

import (
	"github.com/AuroraOps04/mall-go/model/base"
	"gorm.io/gorm"
)

type Brand struct {
	base.BaseModel

	Name                string `json:"name" gorm:"size:64"`
	FirstLetter         string `json:"firstLetter" gorm:"size:8"`
	Sort                int    `json:"sort"`
	FactoryStatus       int8   `json:"factoryStatus" gorm:"size:1"`
	ShowStatus          int8   `json:"showStatus" gorm:"size:1"`
	ProductCount        uint   `json:"productCount"`
	ProductCommentCount uint   `json:"productCommentCount"`
	Logo                string `json:"logo" gorm:"size:255"`
	BigPic              string `json:"bigPic" gorm:"size:255"`
	BrandStory          string `json:"brandStory" gorm:"type:text"`
}

func (p Brand) TableName() string {
	return "pms_brand"
}

// create update hook

func (p Brand) BeforeCreate(db *gorm.DB) error {
	// TODO set firstletter if not exits
	return nil
}

func (p Brand) BeforeUpdate(db *gorm.DB) error {
	// TODO
	//set firstletter if not exits
	// set product.brandName
	return nil
}
