package routes

import (
	"github.com/devfajar/go-bimbingan-skripsi/controllers"
	"github.com/devfajar/go-bimbingan-skripsi/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	adminRoutes := app.Group("api/dashboard", middlewares.AdminOnly)

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.UserDetail)
	app.Post("/logout", controllers.Logout)

	// Admin Only Can Access This Route
	adminRoutes.Get("/students", controllers.ListStudent)

}
