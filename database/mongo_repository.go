package database

import (
	"context"
	"user-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	db *mongo.Database
}

func NewMongoRepository(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		db: db,
	}
}

func (r *MongoRepository) CreateWallet(ctx context.Context, wallet *models.Wallet) error {
	_, err := r.db.Collection("wallets").InsertOne(ctx, wallet)
	return err
}

func (r *MongoRepository) GetWallet(ctx context.Context, id string) (*models.Wallet, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var wallet models.Wallet
	err = r.db.Collection("wallets").FindOne(ctx, bson.M{"_id": objID}).Decode(&wallet)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *MongoRepository) UpdateWallet(ctx context.Context, wallet *models.Wallet) error {
	_, err := r.db.Collection("wallets").UpdateOne(
		ctx,
		bson.M{"_id": wallet.ID},
		bson.M{"$set": wallet},
	)
	return err
}

func (r *MongoRepository) ListWallets(ctx context.Context) ([]*models.Wallet, error) {
	cursor, err := r.db.Collection("wallets").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var wallets []*models.Wallet
	if err = cursor.All(ctx, &wallets); err != nil {
		return nil, err
	}
	return wallets, nil
}
