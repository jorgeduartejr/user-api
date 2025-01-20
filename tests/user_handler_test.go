package tests

import (
    "bytes"
    "net/http"
    "testing"
    "user-api/database"
    "user-api/handlers"

    "github.com/gofiber/fiber/v2"
)

func setupTestApp() *fiber.App {
    app := fiber.New()
    app.Get("/users", handlers.GetUsers)
    app.Post("/users", handlers.CreateUser)
    app.Put("/users/:id", handlers.UpdateUser)
    app.Delete("/users/:id", handlers.DeleteUser)

    database.ConnectTestDB()
    return app
}

func TestCreateUserInvalidData(t *testing.T) {
    err := database.ConnectTestDB()
    if err != nil {
        t.Fatalf("Failed to connect to test database: %v", err)
    }

    originalDB := database.DB
    database.DB = database.TestDB
    defer func() { database.DB = originalDB }()

    app := setupTestApp()

    invalidUserJSON := []byte(`{"name": "Test User"}`) // Missing email field

    req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(invalidUserJSON))
    req.Header.Set("Content-Type", "application/json")
    resp, _ := app.Test(req)

    if resp.StatusCode != http.StatusBadRequest {
        t.Fatalf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
    }
}

