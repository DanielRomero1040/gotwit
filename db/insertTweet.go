package db

import (
	"context"

	"github.com/DanielRomero1040/gotwit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(element models.SaveTweet) (string, bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("tweets")

	registro := bson.M{
		"userid":  element.UserID,
		"message": element.Message,
		"date":    element.Date,
	}

	resutl, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}
	objID, _ := resutl.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
