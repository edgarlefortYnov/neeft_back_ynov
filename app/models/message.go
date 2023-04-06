package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	ConversationId string `gorm:"not null" json:"conversationId"`
	EmitterId      uint   `gorm:"not null" json:"emitterId"`
	Emitter        User   `gorm:"foreignkey:EmitterId"`
	Content        string `gorm:"not null" json:"content"`
}

func Messages(db *gorm.DB) ([]Message, error) {
	var messages []Message

	err := db.Model(&Message{}).Find(&messages).Error

	return messages, err
}

func GetMessage(db *gorm.DB, id uint) (Message, error) {
	var message Message

	err := db.Model(&Message{}).First(&message, id).Error

	return message, err
}

func GetMessagesByConversationId(db *gorm.DB, id string) ([]Message, error) {
	var messages []Message

	err := db.Model(&Message{}).Where("conversation_id = ?", id).Find(&messages).Error

	return messages, err
}

func GetLastFiftyMessagesByConversationId(db *gorm.DB, id string) ([]Message, error) {
	var messages []Message

	err := db.Model(&Message{}).Limit(50).Where("conversation_id = ?", id).Order("updated_at desc").Find(&messages).Error

	return messages, err
}

func GetMessagesByEmitterId(db *gorm.DB, id uint) ([]Message, error) {
	var messages []Message

	err := db.Model(&Message{}).Where("emitter_id = ?", id).Find(&messages).Error

	return messages, err
}

func CreateMessage(db *gorm.DB, message *Message) error {
	err := db.Model(&Message{}).Create(&message).Error

	return err
}

func UpdateMessage(db *gorm.DB, message Message) (Message, error) {
	err := db.Model(&Message{}).Where("id = ?", message.ID).Updates(&message).Error

	return message, err
}

func DeleteMessage(db *gorm.DB, id uint) error {
	err := db.Model(&Message{}).Where("id = ?", id).Delete(&Message{}).Error

	return err
}
