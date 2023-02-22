package tournament

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/models"
	"neeft_back/database"
)

// SerializedTournamentResponse the structure of a tournament that is sent back to the client
type SerializedTournamentResponse struct {
	ID         uint   `json:"id" `
	Name       string `json:"name"`
	Count      uint   `json:"count"`
	Price      uint   `json:"price"`
	Game       string `json:"game"`
	TeamsCount uint   `json:"teamsCount"`
	IsFinished bool   `json:"isFinished"`
	Mode       string `json:"mode"`
}

// CreateTournament creates a new tournament
func CreateTournament(c *fiber.Ctx) error {
	var tournament models.Tournament

	if err := c.BodyParser(&tournament); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// Default values
	tournament.IsFinished = false

	database.Database.Db.Create(&tournament)

	return c.Status(200).JSON(tournament)
}

// GetAllTournaments returns all tournaments
func GetAllTournaments(c *fiber.Ctx) error {
	var allTournament []models.Tournament
	var responseTournaments []models.Tournament

	database.Database.Db.Find(&allTournament)

	for _, tournament := range allTournament {
		responseTournament := tournament
		responseTournaments = append(responseTournaments, responseTournament)
	}

	return c.Status(200).JSON(responseTournaments)
}

// FindTournament find a tournament by its id
func FindTournament(id int, tournament *models.Tournament) error {
	database.Database.Db.Find(&tournament, "id = ?", id)

	if tournament.ID == 0 {
		return errors.New("tournament does not exist")
	}

	return nil
}

// GetTournament returns the tournament with the specified id
func GetTournament(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var tournament models.Tournament

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	if err := FindTournament(id, &tournament); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(tournament)
}

// UpdateTournament updates a tournament
func UpdateTournament(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var tournament models.Tournament

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = FindTournament(id, &tournament)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if tournament.ID == 0 {
		return c.Status(400).JSON("invalid tournament")
	}

	var updateData models.Tournament

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	tournament = updateData

	// Resetting values that clients shouldn't modify
	updateData.IsFinished = tournament.IsFinished

	database.Database.Db.Save(&tournament)

	return c.Status(200).JSON(tournament)

}

// DeleteTournament deletes a tournament
func DeleteTournament(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var tournament models.Tournament

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = FindTournament(id, &tournament)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.Db.Delete(&tournament).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Successfully deleted Tournament")
}
