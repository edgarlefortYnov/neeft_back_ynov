package users

import (
	"errors"
	"neeft_back/app/helper"
	"neeft_back/app/models"
	"neeft_back/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

// SerializedUserResponse User : this is the router for the users not the model of User
type SerializedUserResponse struct {
	ID             uint                  `json:"id"`
	Username       string                `json:"username"`
	Email          string                `json:"email"`
	ProfilePicture string                `json:"profilePicture"`
	Roles          []models.RoleRelation `json:"roles"`
	Status         string                `json:"status"`
}

// NewUserResponse /**
func NewUserResponse(userModel models.User) SerializedUserResponse {
	return SerializedUserResponse{
		ID:       userModel.ID,
		Username: userModel.Username,
		Email:    userModel.Email,
		Roles:    userModel.Roles,
		Status:   "active",
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
	var responseUsers []SerializedUserResponse
	for _, user := range users {
		responseUser := NewUserResponse(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}

// FindUser function to find a user by id in the database
func FindUser(id uint, user *models.User) error {
	database.Database.Db.Preload("RoleRelation").First(&user, id)
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
	user, err := models.GetUserWithRelationship(Db, uint(id))
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// Return the user
	responseUser := NewUserResponse(user)
	return c.Status(200).JSON(responseUser)
}

// UpdateUser function is used to update a user
func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	Db := database.Database.Db
	_, err = models.GetUser(Db, uint(id))
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
	responseUser := NewUserResponse(userUpdate)
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
	_, err = models.GetUser(Db, uint(id))
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
	// Assign role to user by creating a new RoleRelation
	var UserHasRole models.RoleRelation
	UserHasRole.UserID = user.ID
	UserHasRole.RoleID = role.ID
	database.Database.Db.Create(&UserHasRole)

	return c.Status(200).JSON("Role assigned to user successfully")
}

func getUserRoles(c *fiber.Ctx, userID uint) ([]models.Role, error) {
	var userRoles []models.Role
	var userHasRoles []models.RoleRelation
	database.Database.Db.Where("user_id = ?", userID).Find(&userHasRoles)
	for _, userHasRole := range userHasRoles {
		var role models.Role
		database.Database.Db.Where("id = ?", userHasRole.RoleID).Find(&role)
		userRoles = append(userRoles, role)
	}
	return userRoles, nil
}
