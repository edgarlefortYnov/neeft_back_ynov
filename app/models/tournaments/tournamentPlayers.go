package tournaments

import "time"

var StatusNone uint = 0
var StatusPending uint = 1
var StatusJoined uint = 2

type TournamentPlayers struct {
	ID           uint `gorm:"primaryKey" json:"id"`
	TournamentId uint `gorm:"not null" json:"tournamentId"`
	UserId       uint `gorm:"not null" json:"userId"`
	Status       uint `gorm:"not null" json:"status"`
	Created_at   time.Time
	Updated_at   time.Time
	Deleted_at   time.Time
}
