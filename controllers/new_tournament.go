package controllers

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"neeft_back/models"
	"net/http"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type CreateTournamentDTO struct {
	Name string `json:"name"`
	Game  string `json:"game"`
	Price  string `json:"price"`
	NbrTeams  string `json:"nbr_teams"`
}

func NewTournament(c *gin.Context) {
	// Open the database
	db, _ := sql.Open("sqlite3", "./bdd.db")


	var req CreateTournamentDTO

    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusForbidden, gin.H{"message": "Expected JSON format", "code": 403})
        return
    }
	tournamentName := req.Name
	tournamentGame := req.Game
	tournamentPrice := req.Price
	tournamentTeamCount := req.NbrTeams
	teamCount, _ := strconv.Atoi(tournamentTeamCount)
	parsedPrice, err := strconv.Atoi(tournamentPrice)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid price provided", "code": 401})
		db.Close()
		return
	}

	// Check if the tournament name isn't empty
	if len(tournamentName) <= 0 {
		c.JSON(http.StatusForbidden, gin.H{"message": "Tournament name is empty", "code": 401})
		db.Close()
		return
	}

	// Check if the tournament price is higher than 0
	if parsedPrice <= 0 {
		c.JSON(http.StatusForbidden, gin.H{"message": "Tournament price cannot be 0", "code": 401})
		db.Close()
		return
	}

	// Check if the number of team is equals or higher than 2
	if teamCount < 2 {
		c.JSON(http.StatusForbidden, gin.H{"message": "You must have 2 teams at least", "code": 401})
		db.Close()
		return
	}

	// Check if the tournament already exists
	row := db.QueryRow("select * from tournaments where name=? and game=?", tournamentName, tournamentGame)
	tournament := new(models.Tournament)
	err = row.Scan(&tournament.Id, &tournament.Name, &tournament.Count, &tournament.Price, &tournament.Game, &tournament.TeamsCount, &tournament.IsFinished, &tournament.Mode)
	if err == nil && tournament.Name == tournamentName && tournament.IsFinished == 0 {
		c.JSON(http.StatusForbidden, gin.H{"message": "A tournament with the same name already exists and isn't finished", "code": 401})
		db.Close()
		return
	}

	if tournamentGame == "Lol" {
		if (teamCount%2 != 0) {
			c.JSON(http.StatusForbidden, gin.H{"message": "team size must be divisible by 2", "code": 401})
			db.Close()
			return
		} else {
			// Insert an element in a table
			query := "INSERT INTO tournaments(name, count, price, game, nbr_teams, end, mode) VALUES (?, ?, ?, ?, ?, ?, ?)"
			ctx, cancelFunction := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelFunction()
			stmt, err := db.PrepareContext(ctx, query)
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{"message": "Database can't be accessed", "error": err.Error(), "code": 401})
				db.Close()
				return
			}
			defer stmt.Close()

			_, err = stmt.ExecContext(ctx, tournamentName, 1, parsedPrice, tournamentGame, 0, false, "unsupported")
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{"message": "Failed", "error": err.Error(), "code": 401})
				db.Close()
				return
			}
		}
	} else if tournamentGame == "Fortnite" {
		tournamentMode := c.PostForm("mode")

		if tournamentMode != "solo" && tournamentMode != "duo" && tournamentMode != "trio" && tournamentMode != "squad" {
			c.JSON(http.StatusForbidden, gin.H{"message": "Invalid party mode", "code": 401})
			db.Close()
			return
		}

		// Insert an element in a table
		query := "INSERT INTO tournaments(name, count, price, game, nbr_teams, end, mode) VALUES (?, ?, ?, ?, ?, ?, ?)"
		ctx, cancelFunction := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunction()
		stmt, err := db.PrepareContext(ctx, query)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "Database can't be accessed", "error": err.Error(), "code": 401})
			db.Close()
			return
		}
		defer stmt.Close()

		_, err = stmt.ExecContext(ctx, tournamentName, 1, parsedPrice, tournamentGame, 0, false, tournamentMode)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "Failed", "error": err.Error(), "code": 401})
			db.Close()
			return
		}
	}

	// Get the tournament's infos
	getIdRow := db.QueryRow("select * from tournaments where name=? and game=? order by id desc", tournamentName, tournamentGame)
	tournament = new(models.Tournament)
	err = getIdRow.Scan(&tournament.Id, &tournament.Name, &tournament.Count, &tournament.Price, &tournament.Game, &tournament.TeamsCount, &tournament.IsFinished, &tournament.Mode)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error(), "code": 401})
		db.Close()
		return
	}

	c.JSON(http.StatusForbidden, gin.H{"message": "Success", "tournament_id": tournament.Id, "code": 200})

	db.Close()
}