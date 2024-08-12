package seeds

import (
	"fmt"
	"github.com/devfajar/go-bimbingan-skripsi/database"
	"github.com/devfajar/go-bimbingan-skripsi/models"
	"golang.org/x/crypto/bcrypt"
)

func SeedAdminUser() {
	// Check if the admin role exists
	var adminRole models.Role
	if err := database.DB.Where("name = ?", "admin").First(&adminRole).Error; err != nil {
		// Role doesn't exist, create it
		adminRole = models.Role{
			Name: "admin",
		}
		database.DB.Create(&adminRole)
	}

	// Check if the admin user exists
	var adminUser models.User
	if err := database.DB.Where("email = ?", "admin@example.com").First(&adminUser).Error; err != nil {
		// User doesn't exist, create it
		password, _ := bcrypt.GenerateFromPassword([]byte("adminpassword"), bcrypt.DefaultCost)

		adminUser = models.User{
			Name:     "Admin",
			Email:    "admin@example.com",
			Password: password,
			RoleID:   adminRole.Id,
		}
		database.DB.Create(&adminUser)

		fmt.Println("Admin user created with email: admin@example.com and password: adminpassword")
	} else {
		fmt.Println("Admin user already exists.")
	}
}
