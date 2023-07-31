package main

import (
	"datasource/database"
)

type UserTest struct {
	Name string
	Age  int32
}

func main() {
	database.MYSQL_DB.AutoMigrate(&UserTest{})
}
