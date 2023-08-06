package database

import (
	"fmt"
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

type DSNConf struct {
	User      string
	Pwd       string
	HostNPort string
	DbName    string
}

func (conf *DSNConf) GetDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Pwd,
		conf.HostNPort,
		conf.DbName)
}
