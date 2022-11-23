package main

import (
	"neeft_back/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	controllers.InitDatabase()

	r := gin.Default()

	r.GET("/", controllers.Accueil)
	r.POST("/connect", controllers.Connect)
	r.POST("/register", controllers.Register)
	r.POST("/new_team", controllers.NewTeam)
	r.POST("/new_tournament", controllers.NewTournament)

	r.Run()
}