package field

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/AuroraOps04/mall-go/model/base"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestLikeNumberInt64(t *testing.T) {
	var n1 = LikeNumberInt64(1)
	var p1 json.Marshaler = n1
	var p2 json.Unmarshaler = &n1
	var p3 driver.Valuer = n1
	var p4 sql.Scanner = &n1
	fmt.Println(p1, p2, p3, p4)
}

func TestLikeNumberFloat64(t *testing.T) {
	var n1 = LikeNumberFloat64(1)
	var p1 json.Marshaler = n1
	var p2 json.Unmarshaler = &n1
	fmt.Println(p1, p2)
}

type Test struct {
	base.BaseModel
	Status         LikeNumberInt64
	StatusPointer  *LikeNumberInt64
	Float64        LikeNumberFloat64
	Float64Pointer *LikeNumberFloat64
}

func TestInGorm(t *testing.T) {
	dialector := mysql.Open(
		"root:123456@tcp(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local",
	)
	db, _ := gorm.Open(dialector)
	err := db.AutoMigrate(&Test{})
	fmt.Println(err)

	var sp = LikeNumberInt64(2)
	var fp = LikeNumberFloat64(2)
	db.Create(&Test{
		Status:         sp,
		StatusPointer:  &sp,
		Float64Pointer: &fp,
		Float64:        fp,
	})
	var test Test
	db.Model(&Test{}).Where("status = ?", 2).First(&test)
	defer db.Delete(&test)
	bytes, err := json.Marshal(test)
	fmt.Println(string(bytes), err)
}
