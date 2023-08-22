package database

import (
	"datasource/database/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

func ReadConfYamlFile(path *string) *DSNConf {
	conf := new(DSNConf)
	if path == nil {
		conf.User = "root"
		conf.Pwd = "12345678"
		conf.DbName = "dousheng"
		conf.HostNPort = "127.0.0.1:3306"
	} else {
		panic("not impl")
	}
	return conf
}

func (conf *DSNConf) MigrateDB() error {
	dsn := conf.GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Errorf("数据库连接失败")
	}
	return db.AutoMigrate(
		&model.User{},
		&model.Video{},
		&model.Comment{},
		&model.Message{},
		&model.Like{},
	)
}
