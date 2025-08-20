package repository

import (
	"context"
	"time"

	"github.com/appu900/carrental/config"
	"github.com/appu900/carrental/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// var userCollection *mongo.Collection = config.GetCollection("users")

func CreateUser(user model.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	userCollection := config.GetCollection("users")
	return userCollection.InsertOne(ctx, user)
}

func FindUserByPhoneNumber(phoneNumber string) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user model.User
	userCollection := config.GetCollection("users")
	err := userCollection.FindOne(ctx, bson.M{"phonenumber": phoneNumber}).Decode(&user)
	return user, err
}
