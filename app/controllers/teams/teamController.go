package teams

import (
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/models"
	"neeft_back/database"
)

// SerializedTeamResponse the structure of a team that is sent back to the client
type SerializedTeamResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	UserCount   uint   `json:"userCount"`
	IsBanned    bool   `json:"isBanned"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

// NewSerializedTeamResponse  /**
func NewSerializedTeamResponse(userModel models.User, teamModel models.Team) SerializedTeamResponse {
	return SerializedTeamResponse{
		ID:          teamModel.ID,
		UserCount:   teamModel.MaxMembers,
		Name:        teamModel.Name,
		IsBanned:    teamModel.IsBanned,
		Description: teamModel.Description,
		Type:        teamModel.Type,
	}
}

// CreateTeam function to create a team
func CreateTeam(c *fiber.Ctx) error {
	db := database.Database.Db

	var team models.Team
	if err := c.BodyParser(&team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("CreateTeam: " + err.Error())
	}

	user, err := models.GetUser(db, team.OwnerId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("GetUserWithRelationship: " + err.Error())
	}

	// A team can't be banned when it is created
	team.IsBanned = false
	err = models.CreateNewTeam(db, &team)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("CreateNewTeam: " + err.Error())
	}

	return c.Status(200).JSON(NewSerializedTeamResponse(user, team))
}

// GetTeams function to get all teams
func GetTeams(c *fiber.Ctx) error {
	db := database.Database.Db

	teams, _ := models.Teams(db)

	var responseTeams []SerializedTeamResponse
	for _, team := range teams {
		var user models.User
		db.Find(&user, "id = ?", team.OwnerId)
		responseTeam := NewSerializedTeamResponse(user, team)
		responseTeams = append(responseTeams, responseTeam)
	}

	return c.Status(200).JSON(responseTeams)
}

// GetTeamFromOwnerId function to get a team
func GetTeamFromOwnerId(c *fiber.Ctx) error {
	// Get the id from the url
	ownerId, err := c.ParamsInt("id")

	// Check if the id is valid
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}

	// Find the user
	db := database.Database.Db
	team, err := models.GetTeamByOwnerId(db, uint(ownerId))
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// Return the user
	var user models.User
	db.Find(&user, "id = ?", team.OwnerId)

	responseTeam := NewSerializedTeamResponse(user, team)

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
	responseTeam := NewSerializedTeamResponse(user, team)
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
	return c.Status(200).JSON("Successfully deleted user")
}

func UpdateTeam(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	db := database.Database.Db
	err = models.GetTeam(db, uint(id))
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var updateData models.Team
	var user models.User

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	updateData.ID = uint(id)

	team, err := models.UpdateTeam(db, updateData)

	if team.ID == 0 {
		return c.Status(500).JSON("An unknown error occurred")
	}

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	teamResponse := NewSerializedTeamResponse(user, team)
	return c.Status(200).JSON(teamResponse)
}
