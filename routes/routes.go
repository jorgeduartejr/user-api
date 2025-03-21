package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	WalletRoutes(app)
	// Adicione outras rotas conforme necessário
}
