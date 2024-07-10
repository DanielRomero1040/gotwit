package routers

import (
	"encoding/json"
	"strconv"

	"github.com/DanielRomero1040/gotwit/db"
	"github.com/DanielRomero1040/gotwit/models"
	"github.com/aws/aws-lambda-go/events"
)

func GetAllTweetsFollowing(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	resp := models.NewRespApi()

	page := request.QueryStringParameters["page"]
	IDUser := claim.ID.Hex()

	if len(page) < 0 {
		page = "1"
	}
	pagInt, err := strconv.Atoi(page)
	if err != nil {
		return *resp.
			WithMessage("Debe enviar el parametro PAGE como entero mayor a 0 " + err.Error())
	}

	tweets, status := db.GetAllTweetsFollowing(IDUser, int64(pagInt))

	if !status {
		return *resp.WithMessage("Error al leer los Tweets ")
	}

	respJson, err := json.Marshal(tweets)
	if err != nil {
		return *resp.
			WithStatus(500).
			WithMessage("Error al formatear los datos de los tweets en JSON " + err.Error())
	}

	return *resp.
		WithStatus(200).
		WithMessage(string(respJson))
}
