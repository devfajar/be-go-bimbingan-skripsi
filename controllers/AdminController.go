package controllers

import (
	"github.com/devfajar/go-bimbingan-skripsi/database"
	"github.com/devfajar/go-bimbingan-skripsi/models"
	"github.com/gofiber/fiber/v2"
)

func ListStudent(c *fiber.Ctx) error {
	var students []*models.User

	if err := database.DB.Debug().Find(&students).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch students",
		})
	}

	if len(students) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No students found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": students,
	})
}
