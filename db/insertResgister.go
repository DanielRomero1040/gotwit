package db

import (
	"context"

	"github.com/DanielRomero1040/gotwit/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertResgister(user models.User) (string, bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	user.Password, _ = EncrypPassword(user.Password)
	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}
	ObjectID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjectID.String(), true, nil
}
