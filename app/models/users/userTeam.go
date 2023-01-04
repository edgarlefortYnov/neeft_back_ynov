package users

import (
	"neeft_back/app/models/teams"
	"time"
)

type UserTeam struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	UserId        uint       `gorm:"not null" json:"userId"`
	User          User       `gorm:"foreignkey:UserId"`
	TeamId        uint       `gorm:"not null" json:"teamId"`
	Team          teams.Team `gorm:"foreignkey:TeamId"`
	StatusRequest uint       `gorm:"not null" json:"statusRequest"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
