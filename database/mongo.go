package database

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var TestDB *mongo.Database

func Connect() error {
    clientOptions := options.Client().ApplyURI("mongodb://admin:admin123@localhost:27017")
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        return err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        return err
    }

    DB = client.Database("userdb")
    return nil
}

func ConnectTestDB() error {
    clientOptions := options.Client().ApplyURI("mongodb://admin:admin123@localhost:27017")
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        return err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        return err
    }

    TestDB = client.Database("testdb")
    return nil
}