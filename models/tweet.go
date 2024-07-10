package models

type Tweet struct {
	Message string `bson:"Message" json:"Message"`
}
