package model

import "datasource/database"

func MigrateDB() {
	database.MYSQL_DB.AutoMigrate()
}
