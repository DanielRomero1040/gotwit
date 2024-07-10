package routers

import (
	"github.com/DanielRomero1040/gotwit/db"
	"github.com/DanielRomero1040/gotwit/models"
	"github.com/aws/aws-lambda-go/events"
)

func DeleteTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	resp := models.NewRespApi()

	ID := request.QueryStringParameters["id"]

	if len(ID) < 1 {
		return *resp.
			WithMessage("El parametro ID es obligatorio ")
	}

	err := db.DeleteTweet(ID, claim.ID.Hex())
	if err != nil {
		return *resp.
			WithMessage("Ocurrio un error a intentar borrar el tweet " + err.Error())
	}
	return *resp.
		WithMessage("Tweet elimninado con exito").
		WithStatus(200)
}
