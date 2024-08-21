package middlewares

import (
	"strconv"

	"github.com/devfajar/go-bimbingan-skripsi/database"
	"github.com/devfajar/go-bimbingan-skripsi/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AdminOnly(c *fiber.Ctx) error {
	// Retrieve JWT from Header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing Authorization Header",
		})
	}

	// Ensure Bearer Token Format
	tokenString := ""
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Authorization Format",
		})
	}

	// Parse JWT with Claims
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secretKey"), nil
	})
	if err != nil || !token.Valid { // Perbaikan: Memeriksa token.Valid == false
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Token",
		})
	}

	// Extract Claims from Token
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to Parse Token Claims",
		})
	}

	// Get User ID from Claims
	id, _ := strconv.Atoi((*claims)["sub"].(string))
	var user models.User

	// Retrieve User and associated Role from Database
	if err := database.DB.Preload("Role").Where("id = ?", uint(id)).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}

	// Check if the User's Role is Admin
	if user.Role.Name != "Admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Access Denied",
		})
	}

	return c.Next()
}
