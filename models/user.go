package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID 	 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` 
	Name string `json:"name,omitempty" bson:"name"`
	Email string `json:"email,omitempty" bson:"email"`
	Age int `json:"age" bson:"age"`
	Password string `json:"password,omitempty" bson:"password"`

}