package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDBConn() error {

	var dbName = "database/NotificationsDB.db"

	myDB, resultErr := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	DB = myDB

	if resultErr != nil {
		fmt.Println("Failed to connect to database.")
		return resultErr
	} else {
		fmt.Println("Connected successfully to database.")
	}

	return nil
}
