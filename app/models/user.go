package models

/**
 * @Author ANYARONKE Daré Samuel
 */

import "time"

type User struct {
	ID              uint   `gorm:"primaryKey"   json:"id" `
	Username        string `gorm:"varchar(255)" json:"username"  validate:"required,min=3,max=32"`
	FirstName       string `gorm:"varchar(255)" json:"firstName" validate:"required,min=3,max=32"`
	LastName        string `gorm:"varchar(255)" json:"lastName"  validate:"required,min=3,max=32"`
	Email           string `gorm:"varchar(255)" json:"email"     validate:"required,email,min=6,max=32"`
	EmailVerifiedAt time.Time
	Password        string `gorm:"varchar(255)" json:"password"  validate:"required,min=8,max=32"`
	RememberToken   string `gorm:"varchar(100)" json:"rememberToken"`
	BirthDate       string `gorm:"varchar(255)" json:"birthDate"`
	Avatar          string `gorm:"varchar(255)" json:"avatar"`
	LastUserAgent   string `gorm:"varchar(255)" json:"lastUserAgent"`
	IsBan           bool   `gorm:"boolean default:false"      json:"isBan"`
	LastLoginAt     time.Time
	IsSuperAdmin    bool `gorm:"boolean default:false"      json:"isSuperAdmin"`

	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
