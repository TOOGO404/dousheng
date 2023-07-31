package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MYSQL_DB *gorm.DB

func init() {
	dsn := "root:000000@tcp(127.0.0.1:3306)/dousheng?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	MYSQL_DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
