package main

import (
    "user-api/database"
    "user-api/routes"
    "log"

    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    err := database.Connect()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    routes.UserRoutes(app)

    log.Fatal(app.Listen(":3000"))
}