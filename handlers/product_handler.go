package handlers

import (
    "context"
    "user-api/database"
    "user-api/models"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection

func InitProductCollection() {
    productCollection = database.GetCollection("userdb", "products")
}

// GetAllProducts lista todos os produtos
func GetAllProducts(c *fiber.Ctx) error {
    cursor, err := productCollection.Find(context.Background(), bson.M{})
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"message": err.Error()})
    }
    var products []models.Product
    if err := cursor.All(context.Background(), &products); err != nil {
        return c.Status(500).JSON(fiber.Map{"message": err.Error()})
    }
    return c.JSON(products)
}

// GetProduct obtem um produto específico
func GetProduct(c *fiber.Ctx) error {
    id := c.Params("id")
    objID, _ := primitive.ObjectIDFromHex(id)
    var product models.Product
    if err := productCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&product); err != nil {
        return c.Status(500).JSON(fiber.Map{"message": err.Error()})
    }
    return c.JSON(product)
}

// CreateProduct cria um novo produto
func CreateProduct(c *fiber.Ctx) error {
    var product models.Product
    if err := c.BodyParser(&product); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": err.Error()})
    }
    product.ID = primitive.NewObjectID()
    if _, err := productCollection.InsertOne(context.Background(), product); err != nil {
        return c.Status(500).JSON(fiber.Map{"message": err.Error()})
    }
    return c.JSON(product)
}

// UpdateProduct atualiza um produto
func UpdateProduct(c *fiber.Ctx) error {
    id := c.Params("id")
    objID, _ := primitive.ObjectIDFromHex(id)
    var newProduct models.Product
    if err := c.BodyParser(&newProduct); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": err.Error()})
    }
    _, err := productCollection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": newProduct})
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"message": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "Product successfully updated"})
}

// DeleteProduct exclui um produto
func DeleteProduct(c *fiber.Ctx) error {
    id := c.Params("id")
    objID, _ := primitive.ObjectIDFromHex(id)
    _, err := productCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"message": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "Product successfully deleted"})
}