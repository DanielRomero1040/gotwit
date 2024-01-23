package db

import (
	"context"

	"github.com/DanielRomero1040/gotwit/models"
)

func InsertRelation(relation models.Relation) (bool, error) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)

	col := db.Collection("relation")

	_, err := col.InsertOne(ctx, relation)
	if err != nil {
		return false, err
	}
	return true, nil
}
