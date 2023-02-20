package authController

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"neeft_back/app/config"
	usersController "neeft_back/app/controllers/users"
	"neeft_back/app/helper"
	"neeft_back/app/models"
	"neeft_back/database"
	"neeft_back/middleware"
	"neeft_back/utils"
	"time"
)

/**
 * @Author ANYARONKE Dare Samuel
 */

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

// validate user information
func validateUserInformation(userInformation *models.User) []*ErrorResponse {
	var errors []*ErrorResponse
	var validateForm = validator.New()

	// Validate the user information
	err := validateForm.Struct(userInformation)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}

	return errors
}

// EmailExist : Check if the email already exist in the database
func EmailExist(email string) bool {
	var user models.User
	if err := database.Database.Db.Find(&user, "email = ?", email).First(&user).Error; err != nil {
		return false
	}
	return true
}

// Register : Register a new user
func Register(c *fiber.Ctx) error {
	userInformation := new(models.User)

	// Get the user information from the request body
	if err := c.BodyParser(userInformation); err != nil {
		return helper.Return400(c, "Invalid user information")
	}

	// Validate the user information and return the errors if there is any error in the user information provided by the user in the request body (email, username, password, etc...)
	errors := validateUserInformation(userInformation)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Check if the user already exist
	if EmailExist(userInformation.Email) {
		return helper.Return400(c, "Email already exist")
	}

	// Hash the password
	hashedPassword := helper.HashAndSalt([]byte(userInformation.Password))

	// Create the user in the database
	user := models.User{
		Username:     userInformation.Username,
		FirstName:    userInformation.FirstName,
		LastName:     userInformation.LastName,
		Email:        userInformation.Email,
		Password:     hashedPassword,
		IsBan:        false,
		IsSuperAdmin: false,
	}
	database.Database.Db.Create(&user)

	// assign the user to the default role
	err := usersController.AssignRoleToUser(c, 1, uint(models.GUEST_ROLE))
	if err != nil {
		return err
	}
	// Send message to the user that the account has been created successfully
	return helper.Return200(c, "User created successfully")
}

// Login : Login a user and return a token to be used for authentication
func Login(c *fiber.Ctx) error {

	userInformation := new(models.User)

	// Get the user information from the request body
	if err := c.BodyParser(userInformation); err != nil {
		return helper.Return400(c, "Invalid user information")
	}

	// Check if the user exists in the database
	var user models.User
	if err := database.Database.Db.Find(&user, "email = ?", userInformation.Email).First(&user).Error; err != nil {
		return helper.Return401(c, "The email or password is incorrect")
	}

	// Check if the password is correct
	if err := helper.ComparePasswords(user.Password, []byte(userInformation.Password)); !err {
		return helper.Return401(c, "The email or password is incorrect ")
	}

	// Check if the user is banned or not verified yet
	if user.IsBan {
		return helper.Return401(c, "A problem occurred during the connection; please contact the administrator")
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
	})
}

func RefreshLogin(c *fiber.Ctx) error {
	claims := config.JWTClaims{}

	if err := utils.CheckJWT(c, &claims); err != nil {
		return c.Status(401).JSON(err.Error())
	}

	user := models.User{}

	if err := middleware.FindUserByClaim(claims, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// Check if the user has the same user agent as stored
	if user.LastUserAgent != string(c.Request().Header.UserAgent()) {
		return helper.Return401(c, "User Agent has changed")
	}

	// Generate the new access token
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

	// Generate the new refresh token
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
	})
}
