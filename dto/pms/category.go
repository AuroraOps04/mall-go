package pms

import (
	"github.com/AuroraOps04/mall-go/model/base"
	pmsModel "github.com/AuroraOps04/mall-go/model/pms"
)

type UpdateCategoryNavStatusDto struct {
	Ids       []int `json:"ids" form:"ids" binding:"required"`
	NavStatus *int  `json:"navStatus" form:"navStatus" binding:"required"`
}

type UpdateCategoryShowStatusDto struct {
	Ids        []int `json:"ids" binding:"required" form:"ids"`
	ShowStatus *int  `json:"showStatus" form:"showStatus" binding:"required"`
}

type CategoryDto struct {
	pmsModel.ProductCategory
	ProductAttributeIdList []uint `json:"productAttributeIdList"`
}

func (d *CategoryDto) ProductAttributes() (list []*pmsModel.ProductAttribute) {
	for _, u := range d.ProductAttributeIdList {
		list = append(list, &pmsModel.ProductAttribute{
			BaseModel: base.BaseModel{
				ID: u,
			},
		})
	}
	return
}
