package models

import "gorm.io/gorm"

type TeamRegistrationTournament struct {
	gorm.Model
	ID                 uint       `gorm:"primaryKey"   json:"id" `
	TeamId             uint       `gorm:"not null" json:"teamId"`
	Team               Team       `gorm:"foreignkey:TeamId"`
	TournamentId       uint       `gorm:"not null" json:"tournamentId"`
	Tournament         Tournament `gorm:"foreignkey:TournamentId"`
	RegistrationStatus uint       `gorm:"not null" json:"registrationStatus"`
}
