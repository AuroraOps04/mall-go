package model

import (
	"github.com/AuroraOps04/mall-go/model/pms"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	pms.Migrate(db)
}
