package models

import "gorm.io/gorm"

type AddFriend struct {
	gorm.Model
	ID       uint `json:"id" gorm:"primaryKey"`
	UserId   int  `gorm:"not null" json:"userId"`
	User     User `gorm:"foreignkey:UserId"`
	FriendId int  `gorm:"not null" json:"friendId"`
	Friend   User `gorm:"foreignkey:FriendId"`
	IsFriend bool `gorm:"bool" json:"is_friend"`

	Status uint `gorm:"not null" json:"status"` // if true, then the user is a friend and delete in this tab after push to user team table
}
