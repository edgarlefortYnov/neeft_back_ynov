package conversations

import (
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/models"
	"neeft_back/database"
)

func CreateMessage(c *fiber.Ctx) error {
	db := database.Database.Db
	var message models.Message

	if err := c.BodyParser(&message); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("chat.err.wrongInformation")
	}
	user, err := models.GetUser(db, message.EmitterId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetUserWithRelationship: " + err.Error())
	}
	message.Emitter = user
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
	var responseMessages []models.Message

	allMessages, err := models.Messages(db)
	for _, message := range allMessages {
		user, err := models.GetUser(db, message.EmitterId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("GetUserWithRelationship: " + err.Error())
		}
		message.Emitter = user
		responseMessages = append(responseMessages, message)
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("GetMessages: " + err.Error())
	}
	return c.Status(200).JSON(responseMessages)
}

func GetMessage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	db := database.Database.Db
	var message models.Message

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	message, err = models.GetMessage(db, uint(id))
	user, err := models.GetUser(db, message.EmitterId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetUserWithRelationship: " + err.Error())
	}
	message.Emitter = user
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetMessage: " + err.Error())
	}
	return c.Status(200).JSON(message)
}

func GetMessagesByConversationId(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.Database.Db
	var messages []models.Message
	var responseMessages []models.Message

	messages, err := models.GetMessagesByConversationId(db, id)
	for _, message := range messages {
		user, err := models.GetUser(db, message.EmitterId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("GetUserWithRelationship: " + err.Error())
		}
		message.Emitter = user
		responseMessages = append(responseMessages, message)
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetMessagesByConversationId: " + err.Error())
	}
	return c.Status(200).JSON(responseMessages)
}

func GetLastFiftyMessagesByConversationId(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.Database.Db
	var messages []models.Message
	var responseMessages []models.Message

	messages, err := models.GetLastFiftyMessagesByConversationId(db, id)
	for _, message := range messages {
		user, err := models.GetUser(db, message.EmitterId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("GetUserWithRelationship: " + err.Error())
		}
		message.Emitter = user
		responseMessages = append(responseMessages, message)
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetLastFiftyMessagesByConversationId: " + err.Error())
	}
	return c.Status(200).JSON(responseMessages)
}

func GetMessagesByEmitterId(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	db := database.Database.Db
	var messages []models.Message
	var responseMessages []models.Message

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	messages, err = models.GetMessagesByEmitterId(db, uint(id))
	for _, message := range messages {
		user, err := models.GetUser(db, message.EmitterId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("GetUserWithRelationship: " + err.Error())
		}
		message.Emitter = user
		responseMessages = append(responseMessages, message)
	}
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetMessagesByEmitterId: " + err.Error())
	}
	return c.Status(200).JSON(responseMessages)
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
	responseMessage, err := models.GetMessage(db, message.ID)
	responseMessage.Emitter, err = models.GetUser(db, responseMessage.EmitterId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetUserWithRelationship: " + err.Error())
	}
	return c.Status(200).JSON(responseMessage)
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
