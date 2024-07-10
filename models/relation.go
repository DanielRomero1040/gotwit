package models

type Relation struct {
	UserID         string `bson:"userid" json:"userid"`
	UserRelationId string `bson:"userielationid" json:"userielationid"`
}
