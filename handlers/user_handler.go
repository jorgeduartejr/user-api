package handlers

import (
    "context"
    "user-api/database"
    "user-api/models"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/gofiber/fiber/v2"
)

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

func CreateUser(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": err.Error()})
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