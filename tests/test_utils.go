package tests

import (
    "user-api/database"
    "user-api/handlers"
    "user-api/routes"
    "github.com/gofiber/fiber/v2"
)

func SetupApp() *fiber.App {
    app := fiber.New()

    // Conexão com o banco de dados de teste
    err := database.ConnectTestDB()
    if err != nil {
        panic(err)
    }

    // Inicializar a coleção de produtos
    handlers.InitProductCollection("testdb")

    // Inicializar a coleção de carteiras
    handlers.InitWalletCollection("testdb")

    // Configuração das rotas de usuários
    routes.UserRoutes(app)

    // Configuração das rotas de produtos
    routes.ProductRoutes(app)

    // Configuração das rotas de carteiras
    routes.WalletRoutes(app)

    return app
}