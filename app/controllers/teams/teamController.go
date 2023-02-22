package teams

/**
 * @Author: ANYARONKE Daré Samuel
 */

import (
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/models"
	"neeft_back/database"
)

// TeamSerialize  User : this is the router for the users not the model of User
// TeamSerialize serializer
type TeamSerialize struct {
	ID          uint        `json:"id"`
	Name        string      `json:"name"`
	UserCount   uint        `json:"userCount"`
	IsBanned    bool        `json:"isBanned"`
	Description string      `json:"description"`
	OwnerId     models.User `json:"ownerId"`
	Type        string      `json:"type"`
}

// CreateResponseTeam  /**
func CreateResponseTeam(userModel models.User, teamModel models.Team) TeamSerialize {
	return TeamSerialize{
		ID:          teamModel.ID,
		UserCount:   teamModel.MaxMembers,
		Name:        teamModel.Name,
		IsBanned:    teamModel.IsBanned,
		Description: teamModel.Description,
		OwnerId:     models.User{Username: userModel.Username},
		Type:        teamModel.Type,
	}
}

// CreateTeam function to create a team
func CreateTeam(c *fiber.Ctx) error {
	Db := database.Database.Db
	var team models.Team
	if err := c.BodyParser(&team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user, err := models.GetUserWithRelationShip(Db, uint(team.OwnerId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	// create a new team
	team.IsBanned = false
	err = models.CreateNewTeam(Db, team)
	responseTeam := CreateResponseTeam(user, team)
	return c.Status(200).JSON(responseTeam)
}

// GetAllTeam function to get all teams
func GetAllTeam(c *fiber.Ctx) error {

	Db := database.Database.Db
	teams, _ := models.Teams(Db)
	var responseTeams []TeamSerialize
	for _, team := range teams {
		var user models.User
		Db.Find(&user, "id = ?", team.OwnerId)
		responseTeam := CreateResponseTeam(user, team)
		responseTeams = append(responseTeams, responseTeam)
	}
	return c.Status(200).JSON(responseTeams)
}

// GetOnerTeam function to get a team
func GetOnerTeam(c *fiber.Ctx) error {
	// Get the id from the url
	ownerId, err := c.ParamsInt("id")
	// Check if the id is valid
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	// Find the user
	Db := database.Database.Db
	team, err := models.GetTeamByOwnerId(Db, uint(ownerId))
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// Return the user
	var user models.User
	Db.Find(&user, "id = ?", team.OwnerId)
	responseTeam := CreateResponseTeam(user, team)
	return c.Status(200).JSON(responseTeam)
}

// GetTeam function to get a team
func GetTeam(c *fiber.Ctx) error {
	// Get the id from the url
	ownerId, err := c.ParamsInt("id")
	// Check if the id is valid
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	// Find the user
	Db := database.Database.Db
	team, err := models.GetTeamById(Db, uint(ownerId))
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// Return the user
	var user models.User
	Db.Find(&user, "id = ?", team.OwnerId)
	responseTeam := CreateResponseTeam(user, team)
	return c.Status(200).JSON(responseTeam)
}

// DeleteTeam function to delete a team
func DeleteTeam(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	Db := database.Database.Db
	err = models.GetTeam(Db, uint(id))
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// Delete the user
	if err = models.DeleteTeam(Db, uint(id)); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted User")
}

func UpdateTeam(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	Db := database.Database.Db
	err = models.GetTeam(Db, uint(id))
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var updateData models.Team
	var user models.User

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	team, err := models.UpdateTeam(Db, updateData)
	if err != nil {
		return err
	}
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseUser := CreateResponseTeam(user, team)
	return c.Status(200).JSON(responseUser)

}
