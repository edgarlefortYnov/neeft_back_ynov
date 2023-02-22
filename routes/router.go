package routes

import (
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/controllers/authController"
	"neeft_back/app/controllers/users"
	"neeft_back/middleware"
)

func SetupRouters(app *fiber.App) {

	//------------------ Auth ---------------------
	api := app.Group("/api")

	app.Post("/api/register", authController.Register)
	app.Post("/api/login", authController.Login)
	app.Post("/refresh-login", authController.RefreshLogin)

	// Need to be logged in to access the routes below
	//TODO: Add middleware to verify the token after login and refresh login not working
	api.Use(middleware.VerifyJWT)

	api.Post("/refresh-login", authController.RefreshLogin)

	//------------------ Users ------------------
	api.Post("/user", users.CreateUser)
	api.Get("/users", users.GetAllUser)
	api.Get("/user/:id", users.GetUser)
	api.Put("/user/:id", users.UpdateUser)
	api.Delete("/user/:id", users.DeleteUser)

	//------------------ Users Role ------------------
	api.Post("/role", users.CreateRole)
	api.Get("/roles", users.GetRoles)
	////------------------ Users Friend ------------------
	//api.Post("/friend", users.CreateUserFriend)
	//api.Get("/show-friend/:id", users.GetUserFriends)

	////------------------ Teams ------------------
	//<<<<<<< Updated upstream
	//	auth.Post("/team", teams.CreateTeam)
	//	auth.Get("/teams", teams.GetAllTeam)
	//	auth.Get("/team/:id", teams.GetTeam)
	//	auth.Put("/team/:id", teams.UpdateTeam)
	//
	//	////------------------ Tournaments ------------------
	//	auth.Post("/tournament", tournament.CreateTournament)
	//	auth.Get("/tournaments", tournament.GetAllTournament)
	//	auth.Get("/tournament/:id", tournament.GetTournament)
	//	auth.Put("/tournament/:id", tournament.UpdateTournament)
	//	auth.Delete("/tournament/:id", tournament.DeleteTournament)
	//=======
	//	api.Post("/team", teams.CreateTeam)
	//	api.Get("/teams", teams.GetAllTeam)
	//	api.Get("/team/:id", teams.GetTeam)
	//
	//	////------------------ Tournaments ------------------
	//	api.Post("/tournament", tournament.CreateTournament)
	//	api.Get("/tournaments", tournament.GetAllTournament)
	//	api.Get("/tournament/:id", tournament.GetTournament)
	//	api.Delete("/tournament/:id", tournament.DeleteTournament)
	//>>>>>>> Stashed changes

	// Debug
	// api.Get("/jwt/debug", users.JWTDebug)
}
