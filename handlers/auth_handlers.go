package handlers

import (
    "context"
    "time"
    "user-api/database"
    "user-api/models"
    "user-api/utils"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
    }
    user.Password = string(hashedPassword)

    // Save the user to the database
    collection := database.DB.Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err = collection.InsertOne(ctx, user)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to register user"})
    }

    return c.Status(201).JSON(fiber.Map{"message": "User registered successfully"})
}

func Login(c *fiber.Ctx) error {
    var input models.User
    if err := c.BodyParser(&input); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
    }

    // Retrieve the user from the database
    var user models.User
    collection := database.DB.Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err := collection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&user)
    if err == mongo.ErrNoDocuments {
        return c.Status(401).JSON(fiber.Map{"error": "User not found"})
    } else if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve user"})
    }

    // Verify the password
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
    if err != nil {
        return c.Status(401).JSON(fiber.Map{"error": "Invalid password"})
    }

    // Generate JWT
    token, err := utils.GenerateToken(user.ID.Hex())
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
    }

    return c.JSON(fiber.Map{"token": token})
}