package db

import (
	"context"

	"github.com/DanielRomero1040/gotwit/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTweetsFollowing(IDUser string, pagina int64) ([]*models.RespTweetsFollowings, bool) {
	docLimit := int64(20)
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relation")

	// options := options.Find()
	// options.SetLimit(docLimit)
	// options.SetSort(bson.D{{Key: "date", Value: -1}})
	// options.SetSkip((pagina - 1) * docLimit)
	skip := (pagina - 1) * docLimit

	condition := make([]bson.M, 0)
	condition = append(condition, bson.M{"$match": bson.M{"userid": IDUser}})
	condition = append(condition, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	condition = append(condition, bson.M{"$unwind": "$tweet"})
	condition = append(condition, bson.M{"$sort": bson.M{"tweet.date": -1}})
	condition = append(condition, bson.M{"$skip": skip})
	condition = append(condition, bson.M{"$limit": docLimit})

	var tweetListFollowing []*models.RespTweetsFollowings

	cursor, err := col.Aggregate(ctx, condition)
	if err != nil {
		return tweetListFollowing, false
	}

	err = cursor.All(ctx, &tweetListFollowing)
	if err != nil {
		return tweetListFollowing, false
	}
	return tweetListFollowing, true
}
