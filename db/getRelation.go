package db

import (
	"context"

	"github.com/DanielRomero1040/gotwit/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRelation(relation models.Relation) bool {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relation")

	condition := bson.M{
		"userid":         relation.UserID,
		"userrelationid": relation.UserRelationId,
	}

	var result models.Relation
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return false
	}
	return true
}
