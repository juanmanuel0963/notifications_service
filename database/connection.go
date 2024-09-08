package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDBConn() error {

	var dbName = "database/NotificationsDB.db"

	resultDB, resultErr := gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	if resultErr != nil {
		return resultErr
	}

	DB = resultDB

	return nil
}
