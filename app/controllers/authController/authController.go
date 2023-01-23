package authController

/**
 * @Author ANYARONKE Dare Samuel
 */

import (
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/helper"
	"neeft_back/app/models/users"
	"neeft_back/database"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"neeft_back/app/config"
)

// Login : Login a user and return a token to be used for authentication
func Login(c *fiber.Ctx) error {

	userInformation := new(users.User)

	// Get the user information from the request body
	if err := c.BodyParser(userInformation); err != nil {
		return helper.Return400(c, "Invalid user information")
	}

	// Check if the user exists in the database
	var user users.User
	if err := database.Database.Db.Find(&user, "email = ?", userInformation.Email).First(&user).Error; err != nil {
		return helper.Return401(c, "Invalid credentials")
	}

	// Check if the password is correct
	if err := helper.ComparePasswords(user.Password, []byte(userInformation.Password)); !err {
		return helper.Return401(c, "Invalid credentials")
	}

	// Check if the user has the same user agent as stored
	if user.LastUserAgent != string(c.Request().Header.UserAgent()) {
		return helper.Return401(c, "User Agent has changed")
	}

	// Generate the access token
	accessTokenExpiryTime := time.Now().Add(time.Minute * 5)
	accessTokenClaims := &config.JWTClaims{
		Email:            user.Email,
		UserId:           user.ID,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		RegisteredClaims: jwt.RegisteredClaims{Issuer: "neeft", ExpiresAt: jwt.NewNumericDate(accessTokenExpiryTime)},
	}
	accessTokenGen := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessToken, err := accessTokenGen.SignedString(config.JWT_SECRET)
	if err != nil {
		return helper.Return500(c, err.Error())
	}

	// Generate the refresh token
	refreshTokenExpiryTime := time.Now().Add(time.Hour * 24 * 7) // 7 days
	refreshTokenClaims := &config.JWTClaims{
		Email:            user.Email,
		UserId:           user.ID,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		Username:         user.Username,
		RegisteredClaims: jwt.RegisteredClaims{Issuer: "neeft", ExpiresAt: jwt.NewNumericDate(refreshTokenExpiryTime)},
	}
	refreshTokenGen := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshToken, err := refreshTokenGen.SignedString(config.JWT_SECRET)
	if err != nil {
		return helper.Return500(c, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user":          user,
	})
}
