package models

/**
 * @Author ANYARONKE Daré Samuel
 */

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Username        string `gorm:"varchar(255)"  validate:"required,min=3,max=32"`
	ID              uint   `gorm:"primaryKey"`
	FirstName       string `gorm:"varchar(255)"  validate:"required,min=3,max=32"`
	LastName        string `gorm:"varchar(255)"  validate:"required,min=3,max=32"`
	Email           string `gorm:"varchar(255)"  validate:"required,email,min=6,max=32"`
	EmailVerifiedAt time.Time

	Password      string    `gorm:"varchar(255)" validate:"required,min=8,max=32"`
	RememberToken string    `gorm:"varchar(100)"`
	BirthDate     string    `gorm:"varchar(255)"`
	Avatar        string    `gorm:"varchar(255)"`
	LastUserAgent string    `gorm:"varchar(255)"`
	IsBan         bool      `gorm:"boolean default:false"`
	LastLoginAt   time.Time `gorm:"null"`
	IsSuperAdmin  bool      `gorm:"boolean default:false"`

	Created_at time.Time      `gorm:"autoCreateTime"`
	Updated_at time.Time      `gorm:"null "`
	Deleted_at gorm.DeletedAt `gorm:"index"`
}
