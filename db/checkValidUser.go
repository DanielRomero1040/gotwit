package db

import (
	"context"

	"github.com/DanielRomero1040/gotwit/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckValidUser(email string) (models.User, bool, string) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")
	condition := bson.M{"email": email}
	var resultUser models.User

	err := col.FindOne(ctx, condition).Decode(&resultUser)
	ID := resultUser.ID.Hex()
	if err != nil {
		return resultUser, false, ID
	}
	return resultUser, true, ID
}
