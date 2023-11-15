package models

type Sale struct {
	ClotheId   string `bson:"clotheId" json:"clotheId"`
	ClotheName string `bson:"clotheName" json:"clotheName"`
	UserId     string `bson:"userId" json:"userId"`
	Date       string `bson:"date" json:"date"`
}
