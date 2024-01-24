package routers

import (
	"encoding/json"
	"strconv"

	"github.com/DanielRomero1040/gotwit/db"
	"github.com/DanielRomero1040/gotwit/models"
	"github.com/aws/aws-lambda-go/events"
)

func GetUsers(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	resp := models.NewRespApi()

	page := request.QueryStringParameters["page"]
	typeUser := request.QueryStringParameters["type"]
	search := request.QueryStringParameters["search"]
	IDUser := claim.ID.Hex()

	if len(page) < 0 {
		page = "1"
	}
	pagInt, err := strconv.Atoi(page)
	if err != nil {
		return resp.
			WithMessage("Debe enviar el parametro PAGE como entero mayor a 0 " + err.Error())
	}

	users, status := db.GetAllUsers(IDUser, int64(pagInt), search, typeUser)
	if !status {
		return resp.WithMessage("Error al leer los usuarios ")
	}

	respJson, err := json.Marshal(users)
	if err != nil {
		return resp.
			WithStatus(500).
			WithMessage("Error al formatear los datos de los usuarios en JSON " + err.Error())
	}

	return resp.
		WithStatus(200).
		WithMessage(string(respJson))

}
