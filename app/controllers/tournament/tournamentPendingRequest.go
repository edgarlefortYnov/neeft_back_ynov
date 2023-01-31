package tournament

import (
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/config"
	"neeft_back/app/models/tournaments"
	"neeft_back/database"
	"neeft_back/utils"
)

func SendPendingRequestPlayer(c *fiber.Ctx) error {
	tournamentId, err := c.ParamsInt("id")

	claims := config.JWTClaims{}

	if err := utils.CheckJWT(c, &claims); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	var tournament tournaments.Tournament
	database.Database.Db.Find(&tournament, "id = ?", tournamentId)
	if tournament.ID == 0 {
		return c.Status(400).JSON("The tournament with the specified id does not exist")
	}

	var tp tournaments.TournamentPlayers
	database.Database.Db.Find(&tp, "tournament_id = ? and user_id = ?", tournamentId, claims.UserId)
	if tp.ID != 0 {
		return c.Status(400).JSON("This player is already in this tournament")
	}

	tp = tournaments.TournamentPlayers{
		TournamentId: uint(tournamentId),
		UserId:       claims.UserId,
		Status:       tournaments.StatusPending,
	}

	database.Database.Db.Create(&tp)

	if tp.ID == 0 {
		return c.Status(400).JSON("Unknown error occurred")
	}

	return c.Status(200).JSON(fiber.Map{"message": "success", "id": tp.ID})
}

func GetTournamentPlayerPendingRequests(c *fiber.Ctx) error {
	tournamentId, err := c.ParamsInt("id")

	claims := config.JWTClaims{}

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := utils.CheckJWT(c, &claims); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var tournament tournaments.Tournament
	database.Database.Db.Find(&tournament, "id = ?", tournamentId)
	if tournament.ID == 0 {
		return c.Status(400).JSON("The tournament with the specified id does not exist")
	}

	var requests []tournaments.TournamentPlayers
	database.Database.Db.Find(&requests, "tournament_id = ? and status = ?", tournamentId, tournaments.StatusPending)

	return c.Status(200).JSON(requests)
}

func AcceptTournamentPendingRequestPlayer(c *fiber.Ctx) error {
	tournamentId, err := c.ParamsInt("id")
	requestId, err := c.ParamsInt("rid")

	claims := config.JWTClaims{}

	if err := utils.CheckJWT(c, &claims); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	var tournament tournaments.Tournament
	database.Database.Db.Find(&tournament, "id = ?", tournamentId)
	if tournament.ID == 0 {
		return c.Status(400).JSON("The tournament with the specified id does not exist")
	}

	// TODO: Check if this client is an admin of the tournament

	var tp tournaments.TournamentPlayers
	database.Database.Db.Find(&tp, "id = ?", requestId)
	if tp.ID != uint(requestId) {
		return c.Status(400).JSON("Invalid request id")
	}
	if tp.Status != tournaments.StatusPending {
		return c.Status(400).JSON("Request can't be accepted")
	}

	tp.Status = tournaments.StatusJoined
	database.Database.Db.Save(&tp)

	return c.Status(200).JSON("success")
}
