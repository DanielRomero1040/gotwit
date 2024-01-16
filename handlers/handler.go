package handlers

import (
	"context"
	"fmt"
	"slices"

	"github.com/DanielRomero1040/gotwit/jwt"
	"github.com/DanielRomero1040/gotwit/models"
	"github.com/DanielRomero1040/gotwit/routers"
	"github.com/aws/aws-lambda-go/events"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))
	var res = models.NewRespApi()

	isOK, statusCode, msg, _ := checkAuthorization(ctx, request)

	if !isOK {
		res.WithStatus(statusCode).WithMessage(msg)
		return res
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "register":
			return routers.Register(ctx)
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

var freePath = []string{
	"registro", "login", "getAvatar", "getBanner",
}

func checkAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)

	if slices.Contains(freePath, path) {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Token requerido", models.Claim{}
	}

	claim, isOK, msg, err := jwt.TokenProccess(token, ctx.Value(models.Key("jwtsign")).(string))

	if !isOK {
		if err != nil {
			fmt.Println("Error en el token " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error en el token" + msg)
			return false, 401, msg, models.Claim{}
		}
	}

	fmt.Println("Token OK")
	return true, 200, msg, *claim
}
