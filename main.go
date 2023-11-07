package main

import (
	pmsView "github.com/AuroraOps04/mall-go/apps/admin/pms"
	"github.com/AuroraOps04/mall-go/apps/common"
	"github.com/AuroraOps04/mall-go/global"
	"github.com/AuroraOps04/mall-go/model/pms"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dialector := mysql.Open("root:123456@tcp(127.0.0.1)/gormdemo?charset=utf8mb4&parseTime=True&loc=Local")
	//dialector := mysql.Open("reader:123456@tcp(192.168.59.10)/mall?charset=utf8mb4&parseTime=True&loc=Local")
	global.Db, _ = gorm.Open(dialector, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	pms.Migrate(global.Db)
	engine := gin.Default()
	adminGroup := engine.Group("/admin")
	pmsView.Route(adminGroup)
	common.Route(engine)
	engine.Static("/static", "./static")
	engine.Run(":4000")

}
