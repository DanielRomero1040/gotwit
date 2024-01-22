package db

import (
	"context"

	"github.com/DanielRomero1040/gotwit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadTweets(ID string, pagina int64) ([]*models.RespTweets, bool) {
	docLimit := int64(20)
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("tweets")
	var resultado []*models.RespTweets

	condition := bson.M{
		"userid": ID,
	}

	options := options.Find()
	options.SetLimit(docLimit)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((pagina - 1) * docLimit)

	cursor, err := col.Find(ctx, condition, options)

	if err != nil {
		return resultado, false
	}

	for cursor.Next(ctx) {
		var registro models.RespTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, false
		}
		resultado = append(resultado, &registro)
	}

	return resultado, true

}
