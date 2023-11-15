package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Clothe struct {
	Type      string `bson:"type"`
	Name      string `bson:"name"`
	Img_front string `bson:"img_front"`
	Img_back  string `bson:"img_back"`
	Sizes     Sizes  `bson:"sizes"`
	Price     int    `bson:"price"`
}

type ClotheWithId struct {
	Name      string             `json:"name"`
	Type      string             `json:"type"`
	Img_front string             `json:"img_front"`
	Img_back  string             `json:"img_back"`
	Sizes     Sizes              `json:"sizes"`
	Price     int                `json:"price"`
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
}

var ClotheTypes = []string{"pant", "hoodie", "shirt"}

type Sizes struct {
	M   int `json:"m" bson:"m"`
	S   int `json:"s" bson:"s"`
	L   int `json:"l" bson:"l"`
	XL  int `json:"xl" bson:"xl"`
	XL2 int `json:"xl2" bson:"xl2"`
	XL3 int `json:"xl3" bson:"xl3"`
}
