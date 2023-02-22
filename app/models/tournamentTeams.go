package models

import "gorm.io/gorm"

type TournamentTeams struct {
	gorm.Model
	ID           uint       `json:"id" gorm:"primaryKey"`
	TournamentId uint       `gorm:"not null" json:"tournamentId"`
	Tournament   Tournament `gorm:"foreignkey:TournamentId"`
	TeamId       uint       `gorm:"not null" json:"teamId"`
	Team         Team       `gorm:"foreignkey:TeamId"`
	Status       uint       `gorm:"not null" json:"status"`
}
