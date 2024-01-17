package db

import (
	"context"

	"github.com/DanielRomero1040/gotwit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindProfile(id string) (models.User, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")
	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		return profile, err
	}
	return profile, nil
}
