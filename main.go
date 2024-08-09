package main

import (
	"fmt"
	"github.com/devfajar/go-bimbingan-skripsi/database"
	"github.com/devfajar/go-bimbingan-skripsi/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Connect to Database
	_, err := database.ConnectDB()
	if err != nil {
		panic("could not connect to database")
	}

	// Success Condition
	fmt.Println("Connection is successful")

	// start the fiber app
	app := fiber.New()

	// CORS
	app.Use(cors.New(cors.Config{
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Content-Type,Authorization,Accept,Origin,Access-Control-Request-Method,Access-Control-Request-Headers,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Access-Control-Allow-Methods,Access-Control-Expose-Headers,Access-Control-Max-Age,Access-Control-Allow-Credentials",
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:8080",
	}))

	// Routes
	routes.SetUpRoutes(app)

	// Start Server
	err = app.Listen(":8080")
	if err != nil {
		panic("Could not start server")
	}
}
