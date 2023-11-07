package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Clothe struct {
	Type      string `bson:"type"`
	Name      string `bson:"name"`
	Sizes     string `bson:"sizes"`
	Img_front string `bson:"img_front"`
	Img_back  string `bson:"img_back"`
	Price     int    `bson:"price"`
}

type ClotheWithId struct {
	Name      string             `json:"name"`
	Type      string             `json:"type"`
	Sizes     string             `json:"sizes"`
	Img_front string             `json:"img_front"`
	Img_back  string             `json:"img_back"`
	Price     int                `json:"price"`
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
}

var ClotheTypes = []string{"pant", "hoodie", "shirt"}
