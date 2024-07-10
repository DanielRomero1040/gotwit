package routers

import (
	"encoding/json"
	"strconv"

	"github.com/DanielRomero1040/gotwit/db"
	"github.com/DanielRomero1040/gotwit/models"
	"github.com/aws/aws-lambda-go/events"
)

func ReadTweets(request events.APIGatewayProxyRequest) models.RespApi {
	resp := models.NewRespApi()

	ID := request.QueryStringParameters["id"]
	pagina := request.QueryStringParameters["pagina"]
	if len(ID) < 1 {
		return *resp.
			WithMessage("El parametro ID es obligatorio")
	}
	if len(pagina) < 1 {
		pagina = "1"
	}

	pag, err := strconv.Atoi(pagina)
	if err != nil {
		return *resp.
			WithMessage("Debe enviar el parametro pagina como un valor mayor a 0 ")
	}

	tweets, ok := db.ReadTweets(ID, int64(pag))
	if !ok {
		return *resp.
			WithMessage("Error al leer los tweets ")
	}
	respJson, err := json.Marshal(tweets)
	if err != nil {
		return *resp.
			WithMessage("Error al formatear los datos de los usuarios como JSON ").
			WithStatus(500)
	}
	return *resp.
		WithStatus(200).
		WithMessage(string(respJson))
}
