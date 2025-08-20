package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	Role     string `bson:"role" json:"role"`
	PhoneNumber string `bson:"phone_number" json:"phone_number"`
}

