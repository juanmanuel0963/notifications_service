package models

import "gorm.io/gorm"

type Type struct {
	gorm.Model        // Includes fields: ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `json:"name" gorm:"not null"`      // Name field is NOT NULL
	Window     string `json:"window" gorm:"not null"`    // Window field is NOT NULL
	Frequency  int    `json:"frequency" gorm:"not null"` // Frequency field is NOT NULL
}
