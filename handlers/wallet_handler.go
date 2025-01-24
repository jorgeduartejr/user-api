package handlers

import (
    "context"
    "time"
    "user-api/database"
    "user-api/models"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

var walletCollection *mongo.Collection

func InitWalletCollection(databaseName string) {
    walletCollection = database.GetCollection(databaseName, "wallets")
}

// CreateWallet cria uma nova carteira
// @Summary Cria uma nova carteira
// @Description Cria uma nova carteira
// @Tags wallets
// @Accept  json
// @Produce  json
// @Param wallet body models.Wallet true "Wallet"
// @Success 201 {object} models.Wallet
// @Router /wallets [post]
func CreateWallet(c *fiber.Ctx) error {
    var wallet models.Wallet
    if err := c.BodyParser(&wallet); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": err.Error()})
    }
    wallet.ID = primitive.NewObjectID()
    wallet.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
    wallet.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
    if wallet.Balance == nil {
        wallet.Balance = make(map[string]float64)
    }
    if _, err := walletCollection.InsertOne(context.Background(), wallet); err != nil {
        return c.Status(500).JSON(fiber.Map{"message": err.Error()})
    }
    return c.Status(201).JSON(wallet)
}

// GetWallet obtem uma carteira específica
// @Summary Obtem uma carteira específica
// @Description Obtem uma carteira específica
// @Tags wallets
// @Accept  json
// @Produce  json
// @Param id path string true "Wallet ID"
// @Success 200 {object} models.Wallet
// @Router /wallets/{id} [get]
func GetWallet(c *fiber.Ctx) error {
    id := c.Params("id")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Invalid ID"})
    }
    var wallet models.Wallet
    if err := walletCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&wallet); err != nil {
        return c.Status(404).JSON(fiber.Map{"message": "Wallet not found"})
    }
    return c.JSON(wallet)
}

// AddFunds adiciona fundos à carteira
// @Summary Adiciona fundos à carteira
// @Description Adiciona fundos à carteira
// @Tags wallets
// @Accept  json
// @Produce  json
// @Param id path string true "Wallet ID"
// @Param funds body map[string]float64 true "Funds"
// @Success 200 {object} models.Wallet
// @Router /wallets/{id}/funds [put]
func AddFunds(c *fiber.Ctx) error {
    id := c.Params("id")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Invalid ID"})
    }
    var funds map[string]float64
    if err := c.BodyParser(&funds); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Invalid request body"})
    }
    var wallet models.Wallet
    if err := walletCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&wallet); err != nil {
        return c.Status(404).JSON(fiber.Map{"message": "Wallet not found"})
    }
    for currency, amount := range funds {
        wallet.Balance[currency] += amount
    }
    wallet.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
    if _, err := walletCollection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": wallet}); err != nil {
        return c.Status(500).JSON(fiber.Map{"message": err.Error()})
    }
    return c.JSON(wallet)
}

