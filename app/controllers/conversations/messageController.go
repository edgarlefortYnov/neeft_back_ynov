package conversations

import (
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/models"
	"neeft_back/database"
	"strconv"
)

func CreateMessage(c *fiber.Ctx) error {
	db := database.Database.Db
	var message models.Message

	if err := c.BodyParser(&message); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("chat.err.wrongInformation")
	}
	u64, err := strconv.ParseUint(message.EmitterId, 10, 32)
	emitterId := uint(u64)
	user, err := models.GetUser(db, emitterId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetUserWithRelationship: " + err.Error())
	}
	user = user
	u64Conv, err := strconv.ParseUint(message.ConversationId, 10, 32)
	conversationId := uint(u64Conv)
	conversation, err := models.GetConversation(db, conversationId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetConversation: " + err.Error())
	}
	message.Conversation = conversation
	if message.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON("CreateMessage content: " + err.Error())
	}
	err = models.CreateMessage(db, &message)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("CreateMessage: " + err.Error())
	}
	return c.Status(200).JSON(message)
}

func GetMessages(c *fiber.Ctx) error {
	db := database.Database.Db
	var allMessages []models.Message

	allMessages, err := models.Messages(db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("GetMessages: " + err.Error())
	}
	return c.Status(200).JSON(allMessages)
}

func GetMessage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	db := database.Database.Db
	var message models.Message

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	message, err = models.GetMessage(db, uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetMessage: " + err.Error())
	}
	return c.Status(200).JSON(message)
}

func GetMessagesByConversationId(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	db := database.Database.Db
	var messages []models.Message

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	messages, err = models.GetMessagesByConversationId(db, uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetMessagesByConversationId: " + err.Error())
	}
	return c.Status(200).JSON(messages)
}

func GetLastFiftyMessagesByConversationId(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	db := database.Database.Db
	var messages []models.Message

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	messages, err = models.GetLastFiftyMessagesByConversationId(db, uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetLastFiftyMessagesByConversationId: " + err.Error())
	}
	return c.Status(200).JSON(messages)
}

func GetMessagesByEmitterId(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	db := database.Database.Db
	var messages []models.Message

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	messages, err = models.GetMessagesByEmitterId(db, uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetMessagesByEmitterId: " + err.Error())
	}
	return c.Status(200).JSON(messages)
}

func UpdateMessage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	db := database.Database.Db
	var message models.Message

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	message, err = models.GetMessage(db, uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetMessage: " + err.Error())
	}

	var updateData models.Message

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	updateData.ID = uint(id)
	message, err = models.UpdateMessage(db, updateData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("UpdateMessage: " + err.Error())
	}
	return c.Status(200).JSON(message)
}

func DeleteMessage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	db := database.Database.Db
	var message models.Message

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	message, err = models.GetMessage(db, uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetMessage: " + err.Error())
	}
	message = message
	err = models.DeleteMessage(db, uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("DeleteMessage: " + err.Error())
	}
	return c.Status(200).JSON("Successfully deleted message")
}
