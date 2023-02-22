package models

import (
	"gorm.io/gorm"
	"time"
)

type Tournament struct {
	gorm.Model
	Name       string    `gorm:"varchar(255)" json:"name"`
	Price      uint      `gorm:"int" json:"price"`
	GameId     int       `gorm:"not null" json:"gameId"`
	Game       Game      `gorm:"foreignkey:GameId"`
	OwnerId    int       `gorm:"not null" json:"ownerId"`
	Owner      User      `gorm:"foreignkey:OwnerId"`
	TeamsLimit uint      `gorm:"uint" json:"teamsLimit"`
	IsFinished bool      `gorm:"bool" json:"isFinished"`
	Address    string    `gorm:"varchar(255)" json:"address"`
	Mode       string    `gorm:"varchar(255)" json:"mode"`
	StartDate  time.Time `gorm:"datetime nullable" json:"startDate"`
	EndDate    time.Time `gorm:"datetime nullable" json:"endDate"`
}
