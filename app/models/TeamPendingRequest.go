package models

import "gorm.io/gorm"

var TeamRequestStatusPending uint = 1
var TeamRequestStatusAccepted uint = 2

type TeamPendingRequest struct {
	gorm.Model
	UserID uint
	TeamID uint
	Status uint
	User   User `gorm:"foreignkey:UserID"`
	Team   Team `gorm:"foreignkey:TeamID"`
}

func TeamPendingRequests(db *gorm.DB) ([]TeamPendingRequest, error) {
	var requests []TeamPendingRequest

	err := db.Model(&User{}).Find(&requests).Error

	return requests, err
}

func GetPendingRequestsForTeam(db *gorm.DB, teamId uint) ([]TeamPendingRequest, error) {
	var requests []TeamPendingRequest

	err := db.Model(&TeamPendingRequest{}).Where("team_id = ?", teamId).Find(&requests).Error

	return requests, err
}

// CreateNewTeam a new team
func CreateNewTeamPendingRequest(db *gorm.DB, request *TeamPendingRequest) error {
	err := db.Model(&TeamPendingRequest{}).Create(&request).Error

	return err
}
