package conversations

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"neeft_back/app/models"
	"neeft_back/database"
)

type ConversationRequest struct {
	ChatId           string `gorm:"" json:"chatId"`
	ConversationType string `gorm:"not null" json:"conversationType"`
	UserId           []uint `gorm:"not null" json:"userId"`
	Owner            uint   `gorm:"not null" json:"ownership"`
	IsClosed         bool   `gorm:"boolean default:false" json:"isClosed"`
}

func CreateConversation(c *fiber.Ctx) error {
	db := database.Database.Db
	var conversationRequest ConversationRequest
	var conversation models.Conversation
	var responseConversations []models.Conversation
	var chatId = uuid.New().String()
	if err := c.BodyParser(&conversationRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("CreateConversation: " + err.Error())
	}

	for _, userId := range conversationRequest.UserId {
		user, err := models.GetUser(db, userId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("GetUserWithRelationship: " + err.Error())
		}
		user.Password = ""
		if user.ID == conversationRequest.Owner {
			conversation = models.Conversation{ChatId: chatId, ConversationType: conversationRequest.ConversationType, UserId: userId, User: user, Owner: true, IsClosed: false}
		} else {
			conversation = models.Conversation{ChatId: chatId, ConversationType: conversationRequest.ConversationType, UserId: userId, User: user, Owner: false, IsClosed: false}
		}
		err = models.CreateConversation(db, &conversation)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON("CreateConversation: " + err.Error())
		}
		responseConversations = append(responseConversations, conversation)
	}
	return c.Status(200).JSON(responseConversations)
}

func GetConversations(c *fiber.Ctx) error {
	db := database.Database.Db
	allConversations, err := models.Conversations(db)

	var responseConversations []models.Conversation

	for _, conversation := range allConversations {
		user, err := models.GetUser(db, conversation.UserId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("GetUserWithRelationship: " + err.Error())
		}
		user.Password = ""
		conversation.User = user
		responseConversations = append(responseConversations, conversation)
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("GetConversations: " + err.Error())
	}

	return c.Status(200).JSON(responseConversations)
}

func GetConversation(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	db := database.Database.Db
	var conversation models.Conversation

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	conversation, err = models.GetConversation(db, uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetConversation: " + err.Error())
	}
	user, err := models.GetUser(db, conversation.UserId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetUserWithRelationship: " + err.Error())
	}
	user.Password = ""
	conversation.User = user
	return c.Status(200).JSON(conversation)
}

func GetConversationsByUserId(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	db := database.Database.Db
	var conversations []models.Conversation
	var responseConversations []models.Conversation

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	conversations, err = models.GetConversationsByUserId(db, uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetConversation: " + err.Error())
	}
	for _, conversation := range conversations {
		user, err := models.GetUser(db, conversation.UserId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("GetUserWithRelationship: " + err.Error())
		}
		user.Password = ""
		conversation.User = user
		responseConversations = append(responseConversations, conversation)
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetConversation: " + err.Error())
	}
	return c.Status(200).JSON(responseConversations)
}

func GetConversationsByChatId(c *fiber.Ctx) error {
	chatId := c.Params("id")
	db := database.Database.Db
	var conversations []models.Conversation
	var responseConversations []models.Conversation

	conversations, err := models.GetConversationsByChatId(db, chatId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetConversation: " + err.Error())
	}
	for _, conversation := range conversations {
		user, err := models.GetUser(db, conversation.UserId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("GetUserWithRelationship: " + err.Error())
		}
		user.Password = ""
		conversation.User = user
		responseConversations = append(responseConversations, conversation)
	}
	return c.Status(200).JSON(responseConversations)
}

func UpdateConversation(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.Database.Db
	var conversations []models.Conversation
	var responseConversations []models.Conversation

	conversations, err := models.GetConversationsByChatId(db, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetConversation: " + err.Error())
	}

	var updateData models.Conversation

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	updateData.ChatId = id
	for _, conversation := range conversations {
		user, err := models.GetUser(db, conversation.UserId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("GetUserWithRelationship: " + err.Error())
		}
		user.Password = ""
		updateData.User = conversation.User
		conversation, err = models.UpdateConversation(db, updateData)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("UpdateConversation: " + err.Error())
		}
		responseConversations = append(responseConversations, conversation)
	}
	return c.Status(200).JSON(responseConversations)
}

func DeleteConversation(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.Database.Db
	var conversations []models.Conversation

	conversations, err := models.GetConversationsByChatId(db, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetConversation: " + err.Error())
	}
	for _, conversation := range conversations {
		conversation = conversation
		err = models.DeleteConversation(db, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON("DeleteConversation: " + err.Error())
		}
	}
	return c.Status(200).JSON("Successfully deleted conversation")
}
