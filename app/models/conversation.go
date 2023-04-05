package models

import (
	"gorm.io/gorm"
)

type Conversation struct {
	gorm.Model
	ChatId           string `gorm:"not null" json:"chatId"`
	ConversationType string `gorm:"not null" json:"conversationType"`
	UserId           uint   `gorm:"not null" json:"userId"`
	User             User   `gorm:"foreignkey:UserId"`
	Owner            bool   `gorm:"boolean default:false" json:"ownership"`
	IsClosed         bool   `gorm:"boolean default:false" json:"isClosed"`
}

func Conversations(db *gorm.DB) ([]Conversation, error) {
	var conversations []Conversation

	err := db.Model(&Conversation{}).Find(&conversations).Error

	return conversations, err
}

func GetConversation(db *gorm.DB, id uint) (Conversation, error) {
	var conversation Conversation

	err := db.Model(&Conversation{}).First(&conversation, id).Error

	return conversation, err
}

func GetConversationsByUserId(db *gorm.DB, id uint) ([]Conversation, error) {
	var conversations []Conversation

	err := db.Model(&Conversation{}).Where("user_id = ?", id).Find(&conversations).Error

	return conversations, err
}

func GetConversationsByChatId(db *gorm.DB, id string) ([]Conversation, error) {
	var conversations []Conversation

	err := db.Model(&Conversation{}).Where("chat_id = ?", id).Find(&conversations).Error

	return conversations, err
}

func CreateConversation(db *gorm.DB, conversation *Conversation) error {
	err := db.Model(&Conversation{}).Create(&conversation).Error

	return err
}

func UpdateConversation(db *gorm.DB, conversation Conversation) (Conversation, error) {
	err := db.Model(&Conversation{}).Where("chat_id = ?", conversation.ChatId).Updates(&conversation).Error

	return conversation, err
}

func DeleteConversation(db *gorm.DB, id string) error {
	err := db.Model(&Conversation{}).Where("chat_id = ?", id).Delete(&Conversation{}).Error

	return err
}
