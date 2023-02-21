package routes

import (
	"neeft_back/app/controllers/authController"
	"neeft_back/app/controllers/teams"
	"neeft_back/app/controllers/tournament"
	"neeft_back/app/controllers/users"
	"neeft_back/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRouters(app *fiber.App) {

	//------------------ Auth ---------------------
	api := app.Group("/api")
	api.Post("/player/register", authController.Register)
	api.Post("/player/login", authController.Login)
	api.Post("/player/refresh-login", authController.RefreshLogin)

	auth := api.Use(middleware.VerifyJWT)

	//------------------ Users ------------------
	api.Post("/user", users.CreateUser)
	api.Post("/register", users.CreateUser)
	auth.Get("/users", users.GetAllUser)
	auth.Get("/user/:id", users.GetUser)
	auth.Put("/user/:id", users.UpdateUser)
	auth.Delete("/user/:id", users.DeleteUser)

	////------------------ Users Friend ------------------
	//api.Post("/friend", users.CreateUserFriend)
	//api.Get("/show-friend/:id", users.GetUserFriends)

	////------------------ Teams ------------------
	auth.Post("/team", teams.CreateTeam)
	auth.Get("/teams", teams.GetAllTeam)
	auth.Get("/team/:id", teams.GetTeam)
	auth.Put("/team/:id", teams.UpdateTeam)

	////------------------ Tournaments ------------------
	auth.Post("/tournament", tournament.CreateTournament)
	auth.Get("/tournaments", tournament.GetAllTournament)
	auth.Get("/tournament/:id", tournament.GetTournament)
	auth.Put("/tournament/:id", tournament.UpdateTournament)
	auth.Delete("/tournament/:id", tournament.DeleteTournament)

	// Debug
	// api.Get("/jwt/debug", users.JWTDebug)
}
