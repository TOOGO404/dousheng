package dal

import (
	"datasource/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	conf := database.ReadConfYamlFile(nil)
	db, err := gorm.Open(mysql.Open(conf.GetDSN()))
	if err != nil {
		panic(err)
	} else {
		DB = db
	}
	//to do conf database pool here
}
