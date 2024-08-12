package routes

import (
	"github.com/devfajar/go-bimbingan-skripsi/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	//app.Get("/", controllers.Hello)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.UserDetail)
	app.Post("/logout", controllers.Logout)
}
