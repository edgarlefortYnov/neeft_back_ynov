package users

/**
 * @Author ANYARONKE Daré Samuel
 */

import (
	"errors"
	"neeft_back/app/helper"
	"neeft_back/app/models"
	"neeft_back/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

// UserSerialize User : this is the router for the users not the model of User
// UserSerialize serializer
type UserSerialize struct {
	ID          uint                 `json:"id"`
	Username    string               `json:"username"`
	FirstName   string               `json:"first_name"`
	LastName    string               `json:"last_name"`
	Email       string               `json:"email"`
	UserHasRole []models.UserHasRole `json:"user_has_role"`
}

// CreateResponseUser /**
func CreateResponseUser(userModel models.User) UserSerialize {
	return UserSerialize{
		ID:          userModel.ID,
		Username:    userModel.Username,
		FirstName:   userModel.FirstName,
		LastName:    userModel.LastName,
		Email:       userModel.Email,
		UserHasRole: userModel.UserHasRole,
	}
}

// CreateUser Calling create user function to the user model to create a new user
func CreateUser(c *fiber.Ctx) error {
	Db := database.Database.Db
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).JSON(err)
	}
	user.CreatedAt = time.Now()
	user.EmailVerifiedAt = time.Now()
	err = models.Create(Db, user)
	if err != nil {
		return helper.Return500(c, err.Error())
	}
	return helper.Return200(c, "User created successfully")
}

// GetAllUser function to get all users in the database
func GetAllUser(c *fiber.Ctx) error {
	Db := database.Database.Db
	users, _ := models.Users(Db)
	var responseUsers []UserSerialize
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}

// FindUser function to find a user by id in the database
func FindUser(id uint, user *models.User) error {
	database.Database.Db.Preload("UserHasRole").First(&user, id)
	if user.ID == 0 {
		return errors.New("user not found")
	}
	return nil
}

// GetUser function to find a user by id in the database like find function
func GetUser(c *fiber.Ctx) error {
	// Get the id from the url
	id, err := c.ParamsInt("id")
	// Check if the id is valid
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Please ensure that :id is an integer")
	}
	// Find the user
	Db := database.Database.Db
	user, err := models.GetUserWithRelationShip(Db, uint(id))
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// Return the user
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

// UpdateUser function is used to update a user
func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	Db := database.Database.Db
	err = models.GetUser(Db, uint(id))
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// TODO: to review
	var updateData models.User

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	userUpdate, err := models.Update(Db, updateData)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseUser := CreateResponseUser(userUpdate)
	return c.Status(200).JSON(responseUser)
	// TODO: End to review
}

// DeleteUser function to delete a user
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	Db := database.Database.Db
	err = models.GetUser(Db, uint(id))
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// Delete the user
	if err = models.Delete(Db, uint(id)); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted User")
}

// AssignRoleToUser assign role to user by role id and user id
func AssignRoleToUser(c *fiber.Ctx, userID uint, roleID uint) error {

	// Get user and role by id
	var user models.User
	var role models.Role
	// Check if user and role exist
	if err := FindUser(userID, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := FindRole(roleID, &role); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// Assign role to user by creating a new UserHasRole
	var UserHasRole models.UserHasRole
	UserHasRole.UserID = user.ID
	UserHasRole.RoleID = role.ID
	database.Database.Db.Create(&UserHasRole)

	return c.Status(200).JSON("Role assigned to user successfully")
}

func getUserRoles(c *fiber.Ctx, userID uint) ([]models.Role, error) {
	var userRoles []models.Role
	var userHasRoles []models.UserHasRole
	database.Database.Db.Where("user_id = ?", userID).Find(&userHasRoles)
	for _, userHasRole := range userHasRoles {
		var role models.Role
		database.Database.Db.Where("id = ?", userHasRole.RoleID).Find(&role)
		userRoles = append(userRoles, role)
	}
	return userRoles, nil
}
