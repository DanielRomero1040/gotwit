package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RespTweetsFollowings struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userid" json:"userId,omitempty"`
	UserRelationID string             `bson:"userrelationid" json:"userRelationId,omitempty"`
	RespTweets
}
