package routes

import (
	"user-api/handlers"
	"user-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	api := app.Group("/users", middleware.AuthRequired) // Require authentication for all routes in this group

	api.Get("/", handlers.GetUsers)
	api.Put("/:id", handlers.UpdateUser)
	api.Delete("/:id", handlers.DeleteUser)
	// app.Get("/users", handlers.GetUsers)
	// app.Post("/users", handlers.CreateUser)
	// app.Put("/users/:id", handlers.UpdateUser)
	// app.Delete("/users/:id", handlers.DeleteUser)
}