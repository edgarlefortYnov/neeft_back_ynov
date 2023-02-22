package users

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/helper"
	"neeft_back/app/models"
	"neeft_back/database"
)

type UserHasRoleSerialize struct {
	Role models.Role `json:"Role"`
}

// CreateUserHasRoleResponseUser /**
func CreateUserHasRoleResponseUser(userHasRole models.UserHasRole) UserHasRoleSerialize {
	return UserHasRoleSerialize{
		Role: userHasRole.Role,
	}
}

// CreateRole create a new role
func CreateRole(c *fiber.Ctx) error {
	var role models.Role
	if err := c.BodyParser(&role); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&role)

	return helper.Return200(c, "Role created successfully")
}

// GetRoles get all roles
func GetRoles(c *fiber.Ctx) error {
	var roles []models.Role
	database.Database.Db.Find(&roles)
	return c.Status(200).JSON(roles)
}

// FindRole function to find a user by id in the database
func FindRole(id uint, role *models.Role) error {
	database.Database.Db.Find(&role, "id = ?", id)
	if role.ID == 0 {
		return errors.New("role does not exist")
	}
	return nil
}

//func GetUserRoles(c *fiber.Ctx, user_id) error {
//	var roles []models.Role
//	id, err := c.ParamsInt("id")
//	if err != nil {
//		return c.Status(400).JSON(err.Error())
//	}
//	database.Database.Db.Find(&roles, "user_id = ?", id)
//	return c.Status(200).JSON(roles)
//}
