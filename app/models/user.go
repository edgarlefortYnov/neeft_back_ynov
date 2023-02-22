package models

/**
 * @Author ANYARONKE Daré Samuel
 */

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username        string `gorm:"varchar(255)"  validate:"required,min=3,max=32"`
	FirstName       string `gorm:"varchar(255)"  validate:"required,min=3,max=32"`
	LastName        string `gorm:"varchar(255)"  validate:"required,min=3,max=32"`
	Email           string `gorm:"varchar(255)"  validate:"required,email,min=6,max=32"`
	EmailVerifiedAt time.Time

	Password      string         `gorm:"varchar(255)" validate:"required,min=8,max=32"`
	RememberToken string         `gorm:"varchar(100)"`
	BirthDate     string         `gorm:"varchar(255)"`
	Avatar        string         `gorm:"varchar(255)"`
	LastUserAgent string         `gorm:"varchar(255)"`
	IsBan         bool           `gorm:"boolean default:false"`
	LastLoginAt   time.Time      `gorm:"null"`
	IsSuperAdmin  bool           `gorm:"boolean default:false"`
	UserHasRole   []UserHasRole  `gorm:"foreignKey:UserID"`
	Created_at    time.Time      `gorm:"autoCreateTime"`
	Updated_at    time.Time      `gorm:"null "`
	Deleted_at    gorm.DeletedAt `gorm:"null"`
}

// Users relationship of the user model with other models that are related to the user
func Users(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("UserHasRole").Find(&users).Error
	return users, err
}

// GetUserWithRelationShip GetOne relationship of the user model with other models that are related to the user
func GetUserWithRelationShip(db *gorm.DB, id uint) (User, error) {
	var user User
	err := db.Model(&User{}).Preload("UserHasRole").First(&user, id).Error
	return user, err
}

// GetUser GetOne user
func GetUser(db *gorm.DB, id uint) error {
	var user User
	err := db.Model(&User{}).First(&user, id).Error
	return err
}

// Create a new user
func Create(db *gorm.DB, user User) error {
	err := db.Model(&User{}).Create(&user).Error
	return err
}

// Update a user
func Update(db *gorm.DB, user User) (User, error) {
	err := db.Model(&User{}).Where("id = ?", user.ID).Updates(&user).Error
	return user, err
}

// Delete a user
func Delete(db *gorm.DB, id uint) error {
	err := db.Model(&User{}).Where("id = ?", id).Delete(&User{}).Error
	return err
}

//--------------------- SCOPES ---------------------//

// ScopeEmailVerifiedAt scope to get all users that have verified their email
func ScopeEmailVerifiedAt(db *gorm.DB) *gorm.DB {
	return db.Where("email_verified_at IS NOT NULL")
}

// ScopeIsBan scope to get all users that are banned
func ScopeIsBan(db *gorm.DB) *gorm.DB {
	return db.Where("is_ban = ?", true)
}

// ScopeIsSuperAdmin scope to get all users that are super admin
func ScopeIsSuperAdmin(db *gorm.DB) *gorm.DB {
	return db.Where("is_super_admin = ?", true)
}
