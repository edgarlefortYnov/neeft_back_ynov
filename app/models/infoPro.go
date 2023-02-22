package models

import "gorm.io/gorm"

type InfoPro struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `gorm:"varchar(255)" json:"name"`
	TeamId uint   `gorm:"not null" json:"teamId"`
	Team   Team   `gorm:"foreignkey:TeamId"`
}
