package database

import (
	"errors"

	"github.com/juanmanuel0963/notification_service/m/models"
)

func PopulateDB() error {

	// Auto-migrate the tables
	DB.AutoMigrate(&models.Type{}, &models.Notification{})

	// Select * from types table
	var types []models.Type // Use a slice to get multiple records
	typesQuery := DB.Find(&types)

	if typesQuery.Error != nil {
		// Error occurred during the query
		return errors.New("Types table not found. " + typesQuery.Error.Error())
	}

	// Check if the types table is empty
	if typesQuery.RowsAffected == 0 {
		// ---------- Populate types table ----------
		var newTypes = []models.Type{
			{Name: "STATUS", Window: "MINUTE", Frequency: 2},
			{Name: "NEWS", Window: "HOUR", Frequency: 3},
			{Name: "MARKETING", Window: "DAY", Frequency: 1},
		}

		// Create new entries in the types table
		result := DB.Create(&newTypes)

		if result.Error != nil {
			return errors.New("Error populating Types table: " + result.Error.Error())
		}
	}

	return nil
}
