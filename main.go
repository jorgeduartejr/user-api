package main

import (
	"log"
	"user-api/config"
	"user-api/database"
	_ "user-api/docs" // Import necessário para o Swagger
	"user-api/handlers"
	"user-api/repository"
	"user-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // Import do Fiber-Swagger
)

// @title Mini E-commerce API
// @version 1.0
// @description API for a mini e-commerce
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
	// Rota de boas-vindas
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("User API - Bem-vindo à API de Usuários, Produtos e Carteiras!")
	})

	cfg := config.LoadConfig()

	var repo repository.Repository

	if cfg.UseMongoDB {
		// Tenta conectar ao MongoDB
		err := database.Connect()
		if err != nil {
			log.Printf("Warning: Failed to connect to MongoDB: %v", err)
			log.Println("Falling back to in-memory storage")
			repo = repository.NewMemoryRepository()
		} else {
			repo = database.NewMongoRepository(database.DB)
		}
	} else {
		log.Println("Using in-memory storage")
		repo = repository.NewMemoryRepository()
	}

	// Inicializa os handlers com o repositório
	handlers.InitHandlers(repo)

	// Configura as rotas
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
