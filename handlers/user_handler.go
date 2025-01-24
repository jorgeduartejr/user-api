package handlers

import (
    "context"
    "time"
    "user-api/database"
    "user-api/models"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// @title User API
// @version 1.0
// @description This is a sample server for managing users.
// @host localhost:3000
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.User
// @Failure 500 {object} fiber.Map
// @Router /users [get]
func GetUsers(c *fiber.Ctx) error {
    var users []models.User
    collection := database.DB.Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"message": err.Error()})
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var user models.User
        cursor.Decode(&user)
        users = append(users, user)
    }
    return c.JSON(users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 201 {object} models.User
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /users [post]
func CreateUser(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Invalid request body"})
    }

    // Validação dos campos obrigatórios
    if user.Name == "" || user.Email == "" {
        return c.Status(400).JSON(fiber.Map{"message": "Name and Email are required"})
    }

    user.ID = primitive.NewObjectID()
    collection := database.DB.Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := collection.InsertOne(ctx, user)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"message": err.Error()})
    }
    return c.Status(201).JSON(user)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update an existing user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.User true "User"
// @Security BearerAuth
// @Success 200 {object} models.User
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /users/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
    id := c.Params("id")
    userId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"message": err.Error()})
    }

    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": err.Error()})
    }

    collection := database.DB.Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    update := bson.M{"$set": user}
    _, err = collection.UpdateOne(ctx, bson.M{"_id": userId}, update)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"message": err.Error()})
    }
    return c.JSON(user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Security BearerAuth
// @Success 204
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
    id := c.Params("id")
    objectId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"message": err.Error()})
    }

    collection := database.DB.Collection("users")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err = collection.DeleteOne(ctx, bson.M{"_id": objectId})
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"message": err.Error()})
    }
    return c.SendStatus(204)
}
