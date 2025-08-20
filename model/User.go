package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct{
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name"`
	Phonenumber string `bson:"phonenumber"`
	Password string `bson:"password"`
	Role string `bson:"role"`
}