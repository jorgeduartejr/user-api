// Register godoc
// @Summary Registro de usuário
// @Description Registra um novo usuário
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.User true "Detalhes do Usuário"
// @Success 201 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /register [post]
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
// Login godoc
// @Summary Login de usuário
// @Description Faz login de um usuário e retorna um token JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "Credenciais do Usuário"
// @Success 200 {object} fiber.Map
// @Failure 401 {object} fiber.Map
// @Router /login [post]

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