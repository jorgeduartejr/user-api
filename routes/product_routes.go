package routes

import (
	"user-api/handlers"

	"github.com/gofiber/fiber/v2"
)

// ProductRoutes define as rotas para o CRUD de produtos
func ProductRoutes(app *fiber.App) {
	products := app.Group("/products")
	products.Get("/", handlers.GetAllProducts)    // Listar todos os produtos
	products.Get("/:id", handlers.GetProduct)    // Obter um produto específico
	products.Post("/", handlers.CreateProduct)   // Criar um novo produto
	products.Put("/:id", handlers.UpdateProduct) // Atualizar um produto
	products.Delete("/:id", handlers.DeleteProduct) // Excluir um produto
}