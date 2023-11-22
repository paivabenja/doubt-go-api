package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Lastname string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	IsAdmin  bool               `json:"isAdmin,omitempty" bson:"isAdmin,omitempty"`
}

type User struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Lastname string `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	IsAdmin  bool   `json:"isAdmin,omitempty" bson:"isAdmin,omitempty"`
}
