package routes

import (
	"user-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/users", handlers.GetUsers)
	app.Post("/users", handlers.CreateUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.DeleteUser)
}