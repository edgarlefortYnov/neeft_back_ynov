package routes

import (
	"neeft_back/app/controllers/teams"
	"neeft_back/app/controllers/tournament"
	"neeft_back/app/controllers/users"

	"github.com/gofiber/fiber/v2"
)

func SetupRouters(app *fiber.App) {

	//------------------ Auth ---------------------
	api := app.Group("/api")
	//api.Post("/player/login", authController.Login)
	//api.Post("/player/refresh-login", authController.RefreshLogin)
	api.Post("/player/register", users.CreateUser)

	//auth := api.Use(middleware.VerifyJWT)

	//------------------ Users ------------------
	api.Post("/user", users.CreateUser)
	api.Post("/register", users.CreateUser)
	api.Get("/users", users.GetAllUser)
	api.Get("/user/:id", users.GetUser)
	api.Put("/user/:id", users.UpdateUser)
	api.Delete("/user/:id", users.DeleteUser)
	//
	////------------------ Users Friend ------------------
	//api.Post("/friend", users.CreateUserFriend)
	//api.Get("/show-friend/:id", users.GetUserFriends)
	//
	////------------------ Teams ------------------
	api.Post("/team", teams.CreateTeam)
	api.Get("/teams", teams.GetAllTeam)
	api.Get("/team/:id", teams.GetTeam)
	//
	////------------------ Tournaments ------------------
	api.Post("/tournament", tournament.CreateTournament)
	api.Get("/tournaments", tournament.GetAllTournament)
	api.Get("/tournament/:id", tournament.GetTournament)
	api.Delete("/tournament/:id", tournament.DeleteTournament)

	// Debug
	// api.Get("/jwt/debug", users.JWTDebug)
}
