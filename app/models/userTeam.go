package models

import "gorm.io/gorm"

type UserTeam struct {
	gorm.Model
	ID            uint `json:"id" gorm:"primaryKey"`
	UserId        uint `gorm:"not null" json:"userId"`
	User          User `gorm:"foreignkey:UserId"`
	TeamId        uint `gorm:"not null" json:"teamId"`
	Team          Team `gorm:"foreignkey:TeamId"`
	StatusRequest uint `gorm:"not null" json:"statusRequest"`
}
