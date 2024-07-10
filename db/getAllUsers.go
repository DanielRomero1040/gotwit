package db

import (
	"context"
	"fmt"

	"github.com/DanielRomero1040/gotwit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsers(IDUser string, pagina int64, search string, typeUser string) ([]*models.User, bool) {
	docLimit := int64(20)
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")
	var userList []*models.User

	options := options.Find()
	options.SetLimit(docLimit)
	options.SetSkip((pagina - 1) * docLimit)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?!)` + search},
	}

	cursor, err := col.Find(ctx, query, options)
	if err != nil {
		return userList, false
	}

	var add bool

	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println("Decode = " + err.Error())
			return userList, false
		}
		var relation models.Relation
		relation.UserID = IDUser
		relation.UserRelationId = user.ID.Hex()

		add = false

		found := GetRelation(relation)

		if typeUser == "new" && !found {
			add = true
		}
		if typeUser == "following" && found {
			add = true
		}
		if IDUser == relation.UserID {
			add = false
		}

		if add {
			user.Password = ""
			userList = append(userList, &user)
		}
	}
	err = cursor.Err()
	if err != nil {
		fmt.Println(" cursor.Err() = " + err.Error())
		return userList, false
	}
	cursor.Close(ctx)
	return userList, true
}
