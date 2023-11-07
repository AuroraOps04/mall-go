package pms

import (
	"github.com/AuroraOps04/mall-go/model/base"
)

type Album struct {
	base.BaseModel
	Name        string      `json:"name" gorm:"size:64"`
	Pics        []*AlbumPic `json:"pics" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CoverPic    string      `json:"coverPic" gorm:"size:1000"`
	PicCount    uint        `json:"picCount"`
	Sort        int         `json:"sort"`
	Description string      `json:"description" gorm:"size:1000"`
}

func (a Album) TableName() string {
	return "pms_album"
}

type AlbumPic struct {
	base.BaseModel
	Pic     string `json:"pic" gorm:"size:1000"`
	AlbumID uint   `json:"albumId"`
}

func (a AlbumPic) TableName() string {
	return "pms_album_pic"
}
