package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Car struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Model       string             `bson:"model"`
	Brand       string             `bson:"brand"`
	Year        int                `bson:"year"`
	PricePerDay float64            `bson:"price_per_day"`
	Price       float64            `bson:"price"`
	Available   bool               `bson:"available"`
	ImageURL    string             `bson:"image_url"`
	Description string             `bson:"description"`
}
