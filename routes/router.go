package routes

import (
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/controllers/authController"
	"neeft_back/app/controllers/teams"
	"neeft_back/app/controllers/tournament"
	"neeft_back/app/controllers/users"
	"neeft_back/middleware"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Authentication
	app.Post("/api/player/register", authController.Register)
	app.Post("/api/player/login", authController.Login)

	// Following routes will require JWT authentication
	auth := api.Use(middleware.VerifyJWT)

	// Refresh login
	auth.Post("/refresh-login", authController.RefreshLogin)

	// Users management
	auth.Get("/users", users.GetAllUser)
	auth.Get("/user/:id", users.GetUser)
	auth.Put("/user/:id", users.UpdateUser)
	auth.Delete("/user/:id", users.DeleteUser)

	// Users roles
	auth.Post("/role", users.CreateRole)
	auth.Get("/roles", users.GetRoles)

	// Users friends
	auth.Post("/friend", users.CreateUserFriend)
	auth.Get("/show-friend/:id", users.GetUserFriends)

	// Teams management
	auth.Post("/team", teams.CreateTeam)
	auth.Get("/teams", teams.GetTeams)
	auth.Get("/team/:id", teams.GetTeam)
	auth.Put("/team/:id", teams.UpdateTeam)
	auth.Get("/team/:id/requests/pending", teams.GetTeamPendingRequests)
	auth.Post("/team/:id/requests/pending", teams.CreateTeamPendingRequest)

	// Tournaments management
	auth.Post("/tournament", tournament.CreateTournament)
	auth.Get("/tournaments", tournament.GetAllTournaments)
	auth.Get("/tournament/:id", tournament.GetTournament)
	auth.Delete("/tournament/:id", tournament.DeleteTournament)
}
