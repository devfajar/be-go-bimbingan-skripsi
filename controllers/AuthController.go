package controllers

import (
	"fmt"
	"github.com/devfajar/go-bimbingan-skripsi/database"
	"github.com/devfajar/go-bimbingan-skripsi/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
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

func Login(c *fiber.Ctx) error {
	fmt.Println("Login Request")

	// Parse Body Request
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing request body",
		})
	}

	// Check User Exist
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.ID == 0 {
		fmt.Println("User Not Found")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	// Compare passwords
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err != nil {
		fmt.Println("Invalid Password: ", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Password",
		})
	}

	// Generate JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(int(user.ID)),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte("secretKey"))
	if err != nil {
		fmt.Println("Error generating Token: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to Generate Token",
		})
	}

	// Set JWT Token to cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
	}
	c.Cookie(&cookie)

	// Authentication successful
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User successfully logged in",
	})
}

func UserDetail(c *fiber.Ctx) error {
	fmt.Println("User Detail Request")

	// Retrieve JWT From Cookies
	cookie := c.Cookies("jwt")

	// Parse JWT With claims
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretKey"), nil
	})

	// Handle Token
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Token",
		})
	}

	// Extract Claims from token
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to Parse Token",
		})
	}

	// Get User ID from token
	id, _ := strconv.Atoi((*claims)["sub"].(string))
	user := models.User{ID: uint(id)}

	database.DB.Where("id = ?", uint(id)).First(&user)

	return c.JSON(user)
}
