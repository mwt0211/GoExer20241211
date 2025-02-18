package ginAndGrom

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// 初始化数据库
func init() {
	dsn := "root:pass@tcp(192.168.150.129:3306)/exer?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Sprintf("数据库初始化异常，%s", err.Error())
	}
}
