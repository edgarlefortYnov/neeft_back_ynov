package database

/**
 * @Author ANYARONKE Dare Samuel
 */

import (
	"gorm.io/gorm"
	"neeft_back/app/models"
)

// RunMigration : Run the migration to initialize the database
func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		// Users
		&models.User{},
		&models.Role{},
		&models.RoleRelation{},
		&models.Team{},
		&models.InfoPro{},
		&models.UsersTeam{},
		&models.Game{},
		&models.Tournament{},
		&models.TeamRegistrationTournament{},
		&models.TournamentPlayer{},
		&models.Bracket{},
		&models.AddFriend{},
	)
	if err != nil {
		return
	}
}
