package models

/**
 * @Author ANYARONKE Daré Samuel
 */

import (
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name        string `gorm:"varchar(255)"`
	Description string `gorm:"varchar(255)"`
	OwnerId     uint
	Type        string `gorm:"varchar(255)"`
	IsBanned    bool   `gorm:"not null default:false"`
	MaxMembers  uint   `gorm:"uint"`
}

// Teams Get all teams
func Teams(db *gorm.DB) ([]Team, error) {
	var teams []Team
	err := db.Model(&Team{}).Find(&teams).Error
	return teams, err
}

// GetTeam GetOne team
func GetTeam(db *gorm.DB, id uint) error {
	var team Team
	err := db.Model(&Team{}).First(&team, id).Error
	return err
}

// GetTeamById GetOne team by id
func GetTeamById(db *gorm.DB, id uint) (Team, error) {
	var team Team
	err := db.Model(&Team{}).Where("id = ?", id).First(&team).Error
	return team, err
}

// GetTeamByOwnerId GetOne team by owner id
func GetTeamByOwnerId(db *gorm.DB, ownerId uint) (Team, error) {
	var team Team
	err := db.Model(&Team{}).Where("owner_id = ?", ownerId).First(&team).Error
	return team, err
}

// CreateNewTeam a new team
func CreateNewTeam(db *gorm.DB, team Team) error {
	err := db.Model(&Team{}).Create(&team).Error
	return err
}

// UpdateTeam Update a team
func UpdateTeam(db *gorm.DB, team Team) (Team, error) {
	err := db.Model(&Team{}).Where("id = ?", team.ID).Updates(&team).Error
	return team, err
}

// DeleteTeam Delete a team
func DeleteTeam(db *gorm.DB, id uint) error {
	err := db.Model(&Team{}).Where("id = ?", id).Delete(&Team{}).Error
	return err
}

// GetTeamWithRelationShip GetOne relationship of the team model with other models that are related to the team
func GetTeamWithRelationShip(db *gorm.DB, id uint) (Team, error) {
	var team Team
	err := db.Model(&Team{}).Preload("TeamMembers").First(&team, id).Error
	return team, err
}

// GetTeamWithRelationShipByOwnerId GetOne relationship of the team model with other models that are related to the team
func GetTeamWithRelationShipByOwnerId(db *gorm.DB, ownerId uint) (Team, error) {
	var team Team
	err := db.Model(&Team{}).Preload("TeamMembers").Where("owner_id = ?", ownerId).First(&team).Error
	return team, err
}

// DeleteTeamByOwnerId Delete a team by owner id
func DeleteTeamByOwnerId(db *gorm.DB, ownerId uint) error {
	err := db.Model(&Team{}).Where("owner_id = ?", ownerId).Delete(&Team{}).Error
	return err
}

//--------------------- SCOPES ---------------------//

// ScopeTeamIsBanned scope to get all teams that are banned
func ScopeTeamIsBanned(db *gorm.DB) *gorm.DB {
	return db.Where("is_banned = ?", true)
}

// ScopeTeamIsNotBanned scope to get all teams that are not banned
func ScopeTeamIsNotBanned(db *gorm.DB) *gorm.DB {
	return db.Where("is_banned = ?", false)
}

// ScopeTeamIsFull scope to get all teams that are full
func ScopeTeamIsFull(db *gorm.DB) *gorm.DB {
	return db.Where("max_members = ?", 5)
}

// ScopeTeamIsNotFull scope to get all teams that are not full
func ScopeTeamIsNotFull(db *gorm.DB) *gorm.DB {
	return db.Where("max_members = ?", 4)
}

// ScopeTeamIsPrivate scope to get all teams that are private
func ScopeTeamIsPrivate(db *gorm.DB) *gorm.DB {
	return db.Where("type = ?", "private")
}

// ScopeTeamIsPublic scope to get all teams that are public
func ScopeTeamIsPublic(db *gorm.DB) *gorm.DB {
	return db.Where("type = ?", "public")
}

// DeletePublicTeam Delete a public team using ScopeTeamIsPublic scope
func DeletePublicTeam(db *gorm.DB, id uint) error {
	err := db.Scopes(ScopeTeamIsPublic).Where("id = ?", id).Delete(&Team{}).Error
	return err
}
