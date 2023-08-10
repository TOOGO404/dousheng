package main

import (
	"datasource/database"
	"log"
)

func main() {
	conf := database.ReadConfYamlFile(nil)
	if err := conf.MigrateDB(); err != nil {
		log.Fatal("创建表失败\n", err.Error())
	} else {
		log.Println("创建表成功")
	}
}
