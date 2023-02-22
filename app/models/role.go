package models

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	gorm.Model
	Name        string `gorm:"varchar(255)"`
	Description string `gorm:"varchar(255)"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}

// List of roles
// TODO: Add roles to the database; Data to be completed
var SUPER_ADMIN = "super_admin"
var GUEST = "guest"
