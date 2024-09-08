package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	Stamp   string `json:"stamp" gorm:"not null"`                                            // Stamp field is NOT NULL
	User    string `json:"user" gorm:"not null"`                                             // User field is NOT NULL
	Subject string `json:"subject" gorm:"not null"`                                          // Subject field is NOT NULL
	TypeID  int    `gorm:"not null"`                                                         // TypeID field is NOT NULL
	Type    Type   `gorm:"foreignKey:TypeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Enforce foreign key relationship with constraints
}
