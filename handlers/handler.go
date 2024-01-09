package handlers

import (
	"context"
	"fmt"

	"github.com/DanielRomero1040/gotwit/models"
	"github.com/aws/aws-lambda-go/events"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))
	var res = models.NewRespApi()
	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		}
		//
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		}
		//
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		}
		//
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		}
		//
	}

	res.WithMessage("Method Invalid")
	return res
}
