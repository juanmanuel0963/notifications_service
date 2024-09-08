package types

import (
	"errors"

	"github.com/juanmanuel0963/notification_service/m/database"
	"github.com/juanmanuel0963/notification_service/m/models"
	"gorm.io/gorm"
)

func GetTypeById(TypeId int) (models.Type, error) {
	var typeRecord models.Type

	// Query the Type from the database by ID
	result := database.DB.First(&typeRecord, TypeId)

	// Check if the query encountered an error
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Type{}, errors.New("type not found")
	} else if result.Error != nil {
		return models.Type{}, result.Error
	}

	return typeRecord, nil
}
