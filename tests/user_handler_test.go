package tests

import (
    "bytes"
    "encoding/json"
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

    var body map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&body)
    t.Logf("Response: %v", body)

    if resp.StatusCode != http.StatusBadRequest {
        t.Fatalf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
    }
}

func TestGetUsers(t *testing.T) {
    err := database.ConnectTestDB()
    if err != nil {
        t.Fatalf("Failed to connect to test database: %v", err)
    }

    originalDB := database.DB
    database.DB = database.TestDB
    defer func() { database.DB = originalDB }()

    app := setupTestApp()

    req, _ := http.NewRequest("GET", "/users", nil)
    resp, _ := app.Test(req)

    var body map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&body)
    t.Logf("Response: %v", body)

    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
    }
}

func TestCreateUserValidData(t *testing.T) {
    err := database.ConnectTestDB()
    if err != nil {
        t.Fatalf("Failed to connect to test database: %v", err)
    }

    originalDB := database.DB
    database.DB = database.TestDB
    defer func() { database.DB = originalDB }()

    app := setupTestApp()

    validUserJSON := []byte(`{"name": "Test User", "email": "test@example.com"}`)

    req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(validUserJSON))
    req.Header.Set("Content-Type", "application/json")
    resp, _ := app.Test(req)

    var body map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&body)
    t.Logf("Response: %v", body)

    if resp.StatusCode != http.StatusCreated {
        t.Fatalf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
    }
}

func TestUpdateUser(t *testing.T) {
    err := database.ConnectTestDB()
    if err != nil {
        t.Fatalf("Failed to connect to test database: %v", err)
    }

    originalDB := database.DB
    database.DB = database.TestDB
    defer func() { database.DB = originalDB }()

    app := setupTestApp()

    // First, create a user to update
    validUserJSON := []byte(`{"name": "Test User", "email": "test@example.com"}`)
    t.Logf("Creating user with data: %s", validUserJSON)
    req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(validUserJSON))
    req.Header.Set("Content-Type", "application/json")
    resp, _ := app.Test(req)

    var body map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&body)
    t.Logf("Response: %v", body)

    if resp.StatusCode != http.StatusCreated {
        t.Fatalf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
    }

    // Extract the created user's ID
    userId := body["_id"].(string)

    // Update the user
    updatedUserJSON := []byte(`{"name": "Updated User", "email": "updated@example.com"}`)
    req, _ = http.NewRequest("PUT", "/users/"+userId, bytes.NewBuffer(updatedUserJSON))
    req.Header.Set("Content-Type", "application/json")
    resp, _ = app.Test(req)

    json.NewDecoder(resp.Body).Decode(&body)
    t.Logf("Response: %v", body)

    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
    }
}

func TestDeleteUser(t *testing.T) {
    err := database.ConnectTestDB()
    if err != nil {
        t.Fatalf("Failed to connect to test database: %v", err)
    }

    originalDB := database.DB
    database.DB = database.TestDB
    defer func() { database.DB = originalDB }()

    app := setupTestApp()

    // First, create a user to delete
    validUserJSON := []byte(`{"name": "Test User", "email": "test@example.com"}`)
    req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(validUserJSON))
    req.Header.Set("Content-Type", "application/json")
    resp, _ := app.Test(req)

    var body map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&body)
    t.Logf("Response: %v", body)

    if resp.StatusCode != http.StatusCreated {
        t.Fatalf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
    }

    // Extract the created user's ID
    userId := body["_id"].(string)

    // Delete the user
    req, _ = http.NewRequest("DELETE", "/users/"+userId, nil)
    resp, _ = app.Test(req)

    json.NewDecoder(resp.Body).Decode(&body)
    t.Logf("Response: %v", body)

    if resp.StatusCode != http.StatusNoContent {
        t.Fatalf("Expected status code %d, got %d", http.StatusNoContent, resp.StatusCode)
    }
}
