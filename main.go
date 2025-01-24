package main

import (
    "log"
    "user-api/database"
    "user-api/handlers"
    "user-api/routes"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/swagger" // Import do Fiber-Swagger
    _ "user-api/docs"            // Import necessário para o Swagger
)

// @title User API
// @version 1.0
// @description API for user management
// @contact.name Jorge Duarte
// @contact.email jorge.duarte@example.com
// @license.name MIT
// @host localhost:3000
// @BasePath /
func main() {
    app := fiber.New()

    // Rota para acessar o Swagger
    app.Get("/swagger/*", swagger.HandlerDefault)

    // Testando outra rota
    app.Get("/test", func(c *fiber.Ctx) error {
        return c.SendString("API funcionando!")
    })

    // Conexão com o banco de dados
    err := database.Connect()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    // Inicializar a coleção de produtos
    handlers.InitProductCollection("userdb")

    // Configuração das rotas de usuários
    routes.UserRoutes(app)

    // Configuração das rotas de produtos
    routes.ProductRoutes(app)

    // Inicializar o servidor
    log.Fatal(app.Listen(":3000"))
}

