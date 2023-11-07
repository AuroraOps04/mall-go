package pms

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&Album{},
		&AlbumPic{},
		&Brand{},
		&Comment{},
		&CommentReply{},
		&FeightTemplate{},
		&MemberPrice{},
		&Product{},
		&ProductAttribute{},
		&ProductAttributeCategory{},
		&ProductAttributeValue{},
		&ProductCategory{},
		&ProductFullReduction{},
		&ProductLadder{},
		&ProductOperateLog{},
		&ProductVertifyRecord{},
		&SkuStock{},
	)
	if err != nil {
		panic(err)
	}
}
