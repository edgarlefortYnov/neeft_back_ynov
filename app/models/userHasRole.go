package models

import (
	"time"
)

type RoleRelation struct {
	RoleID uint
	Role   Role `gorm:"foreignkey:RoleID"`
	UserID uint

	Created_at time.Time
	Updated_at time.Time
}
