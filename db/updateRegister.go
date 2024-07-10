package db

import (
	"context"
	"reflect"
	"strings"

	"github.com/DanielRomero1040/gotwit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateRegister(user models.User, id string) (bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	registro := make(map[string]any)
	userElements := reflect.ValueOf(user).MapRange()
	for userElements.Next() {
		if strings.Contains(userElements.Key().String(), "ID") {
			continue
		}
		if strings.Contains(userElements.Key().String(), "Birthday") {
			registro[strings.ToLower(userElements.Key().String())] = user.Birthday
			continue
		}
		if len(userElements.Value().String()) > 0 {
			registro[strings.ToLower(userElements.Key().String())] = userElements.Value().String()
		}
	}
	uptdString := bson.M{
		"$set": registro,
	}
	objID, _ := primitive.ObjectIDFromHex(id)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, uptdString)
	if err != nil {
		return false, err
	}
	return true, nil
}
