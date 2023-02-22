package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string `gorm:"varchar(255)"`
	Description string `gorm:"varchar(255)"`
}

// List of roles
// TODO: Add roles to the database; Data to be completed

var SUPER_ADMIN = "super_admin"
var GUEST = "guest"
