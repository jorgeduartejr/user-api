package models


import "go.mongodb.org/mongo-driver/bson/primitive"

// Wallet represents the model of a wallet


type Wallet struct {

    ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`

    Balance   map[string]float64  `bson:"balance" json:"balance"`

    CreatedAt primitive.DateTime  `bson:"created_at" json:"created_at"`

    UpdatedAt primitive.DateTime  `bson:"updated_at" json:"updated_at"`

}
