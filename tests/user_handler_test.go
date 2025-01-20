// filepath: /home/jorgeduarte/user-api/tests/user_handler_test.go
package tests

import (
    "net/http"
    "testing"
    "user-api/database"
    "user-api/handlers"

    "github.com/gofiber/fiber/v2"
)

func TestGetUsers(t *testing.T) {
    // Conectar ao banco de dados de teste
    err := database.ConnectTestDB()
    if err != nil {
        t.Fatalf("Failed to connect to test database: %v", err)
    }

    // Substituir a conexão do banco de dados pela de teste
    originalDB := database.DB
    database.DB = database.TestDB
    defer func() { database.DB = originalDB }()

    app := fiber.New()
    app.Get("/users", handlers.GetUsers)

    req, _ := http.NewRequest("GET", "/users", nil)
    resp, _ := app.Test(req)

    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
    }
}