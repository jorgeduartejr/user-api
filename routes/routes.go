package routes

import (
	"user-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Health check route
	app.Get("/health", handlers.HealthCheck)

	// Outras rotas
	WalletRoutes(app)
	// Adicione outras rotas conforme necessário
}
