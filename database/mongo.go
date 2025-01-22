package database

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var TestDB *mongo.Database

func Connect() error {
    clientOptions := options.Client().ApplyURI("mongodb://admin:admin123@localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        return err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Ping(ctx, nil)
    if err != nil {
        return err
    }

    DB = client.Database("userdb")
    log.Println("Connected to MongoDB!")
    return nil
}

func ConnectTestDB() error {
    clientOptions := options.Client().ApplyURI("mongodb://admin:admin123@localhost:27017")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        return err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Ping(ctx, nil)
    if err != nil {
        return err
    }

    TestDB = client.Database("testdb")
    log.Println("Connected to Test MongoDB!")
    return nil
}

func GetCollection(databaseName, collectionName string) *mongo.Collection {
    if databaseName == "testdb" {
        return TestDB.Collection(collectionName)
    }
    return DB.Collection(collectionName)
}