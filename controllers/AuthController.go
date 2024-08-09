package controllers

import (
	"fmt"
	"github.com/devfajar/go-bimbingan-skripsi/database"
	"github.com/devfajar/go-bimbingan-skripsi/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	fmt.Println("Registration Request")

	// Parse Request Body
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing request body",
		})
	}

	// Check Email Already Registered
	var existingUser models.User
	if err := database.DB.Where("email = ?", data["email"]).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email already exists",
		})
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to Hash Password",
		})
	}

	// Create New User
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: hashedPassword,
	}

	// insert to database
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to Create User",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User successfully registered",
	})
}
