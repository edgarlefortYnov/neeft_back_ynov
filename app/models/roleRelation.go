package models

type RoleRelation struct {
	RoleID uint
	Role   Role `gorm:"foreignkey:RoleID"`
	UserID uint
}
