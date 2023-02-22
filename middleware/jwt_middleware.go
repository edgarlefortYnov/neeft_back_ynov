package middleware

import (
	"errors"
	"neeft_back/app/config"
	users "neeft_back/app/models"
	"neeft_back/database"
	"neeft_back/utils"

	"github.com/gofiber/fiber/v2"
)

func FindUserByClaim(claims config.JWTClaims, user *users.User) error {
	database.Database.Db.Find(&user, "id = ?", claims.UserId)

	if user.ID == 0 {
		return errors.New("user does not exist")
	}

	return nil
}

func VerifyJWT(c *fiber.Ctx) error {
	claims := config.JWTClaims{}

	if err := utils.CheckJWT(c, &claims); err != nil {
		return c.Status(401).JSON(err.Error())
	}

	user := users.User{}

	if err := FindUserByClaim(claims, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Next()
}
