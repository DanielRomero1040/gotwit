package db

import (
	"context"
	"fmt"

	"github.com/DanielRomero1040/gotwit/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoCN      *mongo.Client
	DatabaseName string
)

func ConectionDB(ctx context.Context) error {
	user := ctx.Value(models.Key("user")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	connString := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)

	var clientOptions = options.Client().ApplyURI(connString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexion exitosa con la DB")
	MongoCN = client
	DatabaseName = ctx.Value(models.Key("database")).(string)
	return nil
}

func DBConnected() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}
