package database

import (
	"gorm.io/gorm"
	"log"
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
		&models.TeamPendingRequest{},
	)

	if err != nil {
		log.Fatalln("Failed to run migrations : " + err.Error())
		return
	}
}
