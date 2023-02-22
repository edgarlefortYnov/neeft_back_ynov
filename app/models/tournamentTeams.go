package models

import (
	"gorm.io/gorm"
	"time"
)

type TournamentTeams struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	TournamentId uint       `gorm:"not null" json:"tournamentId"`
	Tournament   Tournament `gorm:"foreignkey:TournamentId"`
	TeamId       uint       `gorm:"not null" json:"teamId"`
	Team         Team       `gorm:"foreignkey:TeamId"`
	Status       uint       `gorm:"not null" json:"status"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}

// GetTournamentById GetOne tournament
func GetTournamentById(db *gorm.DB, id uint) (Tournament, error) {
	var tournament Tournament
	err := db.Model(&Tournament{}).First(&tournament, id).Error
	return tournament, err
}

// GetTournamentTeam GetOne tournament
func GetTournamentTeam(db *gorm.DB, id uint) error {
	var tournament TournamentTeams
	err := db.Model(&TournamentTeams{}).First(&tournament, id).Error
	return err
}

// DeleteTeamInTournament Delete a team in tournament
func DeleteTeamInTournament(db *gorm.DB, tournamentId uint, teamId uint) error {
	err := db.
		Model(&TournamentTeams{}).
		Where("tournament_id = ? AND team_id = ?", tournamentId, teamId).
		Delete(&TournamentTeams{}).Error
	return err
}
