package middleware

import (
    "user-api/utils"

    "github.com/gofiber/fiber/v2"
)

func AuthRequired(c *fiber.Ctx) error {
    tokenString := c.Get("Authorization")
    if tokenString == "" {
        return c.Status(401).JSON(fiber.Map{"error": "No token provided"})
    }

    _, err := utils.VerifyToken(tokenString)
    if err != nil {
        return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired token"})
    }

    return c.Next()
}