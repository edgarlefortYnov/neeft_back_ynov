package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"   json:"id" `
	Name        string `gorm:"varchar(255)" json:"name"`
	Description string `gorm:"varchar(255)" json:"description"`
	PlayerLimit uint   `gorm:"uint" json:"playerLimit"`
}
