package controllers

import (
	"github.com/devfajar/go-bimbingan-skripsi/database"
	"github.com/devfajar/go-bimbingan-skripsi/models"
	"github.com/gofiber/fiber/v2"
)

func ListRoles(c *fiber.Ctx) error {
	var roles []models.Role

	if err := database.DB.Find(&roles).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to list roles",
		})
	}

	if len(roles) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No roles found",
		})
	}

	return c.JSON(fiber.Map{
		"data": roles,
	})
}

func AddRole(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse request body",
		})
	}

	// Validate role name
	roleName, exists := data["name"]
	if !exists || roleName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Role name is Required",
		})
	}

	// Check if the role already exists in the database
	var existingRole models.Role
	if err := database.DB.Where("name = ?", roleName).First(&existingRole).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Role already exists",
		})
	}

	// Create New Role
	newRole := models.Role{
		Name: roleName,
	}

	// Save New Role
	if err := database.DB.Create(&newRole).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot create role",
		})
	}

	// success Message
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully created role",
	})
}

func UpdateRole(c *fiber.Ctx) error {
	id := c.Params("id")

	// Parsing data from body
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse request body",
		})
	}

	// Check Body if Name not insert
	roleName, exists := data["name"]
	if !exists || roleName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Role name is Required",
		})
	}

	// Get Role By Id
	var role models.Role
	if err := database.DB.Where("id = ?", id).First(&role).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Role not found",
		})
	}

	// updating role
	role.Name = roleName
	if err := database.DB.Save(&role).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot update role",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully updated role",
		"data":    role,
	})
}

func DeleteRole(c *fiber.Ctx) error {
	// Get ID From Params
	id := c.Params("id")

	// Get Role By Id
	var role models.Role
	if err := database.DB.Where("id = ?", id).First(&role).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Role not found",
		})
	}

	// Delete Role
	if err := database.DB.Delete(&role).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot delete role",
		})
	}

	// Success
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully deleted role",
	})
}
