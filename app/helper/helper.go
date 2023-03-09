package helper

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"neeft_back/app/models"
	"neeft_back/database"
	"regexp"
)

//---------------------Auth ---------------------//

// UsernameExist check if username exist
func UsernameExist(username string) bool {
	var user models.User

	if err := database.Database.Db.Find(&user, "username = ?", username).First(&user).Error; err != nil {
		return false
	}

	return true
}

// FirstNameOrLastNameExists check if the name and surname combine exist in the database
func FirstNameOrLastNameExists(firstname string, lastname string) bool {
	var user models.User

	if err := database.Database.Db.Find(&user, "first_name = ? AND last_name = ?", firstname, lastname).First(&user).Error; err != nil {
		return false
	}

	return true
}

// EmailExist : Check if the email already exist in the database
func EmailExist(email string) bool {
	var user models.User

	if err := database.Database.Db.Find(&user, "email = ?", email).First(&user).Error; err != nil {
		return false
	}

	return true
}

// UserExist check if the user exist
func UserExist(email string, username string, firstname string, lastname string) bool {
	return EmailExist(email) || UsernameExist(username)
}

// HandleError is a helper function to handle errors
func HandleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// HashAndSalt is a helper function to hash and salt a password
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	HandleError(err)

	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)

	if err != nil {
		return false
	}

	return true
}

func CheckEmail(email string) (bool, error) {
	if len(email) < 3 && len(email) > 254 {
		return false, errors.New("invalid email size")
	}
	if m, _ := regexp.MatchString("^[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}$", email); !m {
		return false, errors.New("invalid email")
	}
	return true, nil
}

//---------------------ERRORS---------------------//

func Return200(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": message})
}

func Return400(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": message})
}

func Return401(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": message})
}

func Return403(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": message})
}

func Return404(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": message})
}

func Return500(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": message})
}

func Return501(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"message": message})
}

func Return503(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"message": message})
}

func Return504(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusGatewayTimeout).JSON(fiber.Map{"message": message})
}
