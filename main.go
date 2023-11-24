package main

import (
	pmsView "github.com/AuroraOps04/mall-go/apps/admin/pms"
	"github.com/AuroraOps04/mall-go/apps/common"
	_ "github.com/AuroraOps04/mall-go/docs"
	"github.com/AuroraOps04/mall-go/global"
	"github.com/AuroraOps04/mall-go/model/pms"

	// "github.com/AuroraOps04/mall-go/model/pms"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// main entry function
//
//	@Host						localhost:4000
//	@BasePath					/api
//	@Title						mall api based on go
//	@Description				based on gin gorm
//	@Version					v1.0
//	@contact.email				5239604@qq.com
//
//	@contact.name				AuroraOps04
//	@contact.url				https://auroraops04.com
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	dialector := mysql.Open(
		"root:123456@tcp(10.0.0.10)/mall_go?charset=utf8mb4&parseTime=True&loc=Local",
	)
	//dialector := mysql.Open("reader:123456@tcp(192.168.59.10)/mall?charset=utf8mb4&parseTime=True&loc=Local")
	global.Db, _ = gorm.Open(dialector, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	pms.Migrate(global.Db)

	engine := gin.Default()

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	adminGrouP := engine.Group("/api")
	pmsView.Route(adminGrouP)
	common.Route(adminGrouP)
	engine.Static("/static", "./static")
	engine.Run(":4000")

}
