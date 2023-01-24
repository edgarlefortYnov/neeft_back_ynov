package tournament

/**
 * @Author ANYARONKE
 */

import (
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/config"
	teams2 "neeft_back/app/controllers/teams"
	"neeft_back/app/models/teams"
	"neeft_back/app/models/tournaments"
	"neeft_back/app/models/users"
	"neeft_back/database"
	"neeft_back/middleware"
	"neeft_back/utils"
)

type TournamentTeamsRequest struct {
	TeamId uint `json:"teamId"`
	Status uint `json:"status"`
}

// SendPendingRequest is used to send a pending request to join a tournament by a team
func SendPendingRequest(c *fiber.Ctx) error {
	claims := config.JWTClaims{}

	if err := utils.CheckJWT(c, &claims); err != nil {
		return c.Status(401).JSON(err.Error())
	}

	user := users.User{}

	if err := middleware.FindUserByClaim(claims, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var request TournamentTeamsRequest
	tournamentId, _ := c.ParamsInt("id")

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var tournament tournaments.Tournament
	err := FindTournament(tournamentId, &tournament)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var team teams.Team
	err = teams2.FindTeam(int(request.TeamId), &team)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// Check if the team has already sent a pending request
	var potentialRequest tournaments.TournamentTeams

	if err := database.Database.Db.Where("team_id = ?", request.TeamId).First(&potentialRequest).Error; err == nil {
		if potentialRequest.Status == tournaments.StatusPending {
			return c.Status(400).JSON("This team has already sent a pending request")
		}
	}

	// Check if the user is the team owner
	if uint(team.UserId) != user.ID {
		return c.Status(400).JSON("You are not the owner of the team")
	}

	finalItem := tournaments.TournamentTeams{
		TournamentId: uint(tournamentId),
		TeamId:       request.TeamId,
		Status:       0,
		Tournament:   tournament,
		Team:         team,
	}

	database.Database.Db.Create(&finalItem)
	c.Status(200)
	return nil
}
