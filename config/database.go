package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Mongodb connection failed", err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Mongdb Ping failed", err)
	}
	DB = client.Database(dbName)
	fmt.Println("Connected to mongodb")
}

func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}
