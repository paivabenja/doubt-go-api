package models

type Sale struct {
	ClotheId   string      `bson:"clotheId" json:"clotheId"`
	ClotheName string      `bson:"clotheName" json:"clotheName"`
	UserId     string      `bson:"userId" json:"userId"`
	Date       string      `bson:"date" json:"date"`
	Address    AddressType `bson:"address" json:"address"`
}

type AddressType struct {
	StreetName   string `bson:"streetName" json:"streetName"`
	ExtraData    string `bson:"extraData" json:"extraData"`
	PostalCode   int    `bson:"postalCode" json:"postalCode"`
	StreetNumber int    `bson:"streetNumber" json:"streetNumber"`
}
