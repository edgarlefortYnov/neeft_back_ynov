package main

/**
²* @Author: Neeft, ANYARONKE Daré Samuel
*/
import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"neeft_back/database"
	"neeft_back/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	database.ConnectDb() // Initialize the database if it does not exist (it is created automatically the tables thanks to the migration)

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	routes.SetupRouters(app)
	log.Fatal(app.Listen(":" + port))
}
