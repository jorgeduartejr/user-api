package routes

import (
	"github.com/gofiber/fiber/v2"
	"user-api/handlers"
)

func WalletRoutes(app *fiber.App) {
	app.Post("/wallets", handlers.CreateWallet)
	app.Get("/wallets/:id", handlers.GetWallet)
	app.Put("/wallets/:id", handlers.AddFunds)
}