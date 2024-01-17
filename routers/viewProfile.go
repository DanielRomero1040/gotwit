package routers

import (
	"encoding/json"
	"fmt"

	"github.com/DanielRomero1040/gotwit/db"
	"github.com/DanielRomero1040/gotwit/models"
	"github.com/aws/aws-lambda-go/events"
)

func ViewProfile(request events.APIGatewayProxyRequest) models.RespApi {
	resp := models.NewRespApi()

	fmt.Println("viewProfile")
	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		return resp.WithMessage("El parametro ID es obligatorio ")
	}

	profile, err := db.FindProfile(ID)
	if err != nil {
		return resp.WithMessage("Ocurrio un error al interntar buscar el registro " + err.Error())
	}
	respJson, err := json.Marshal(profile)
	if err != nil {
		return resp.
			WithStatus(500).
			WithMessage("Error al formatear los datos de los usuarios como Json " + err.Error())
	}

	return resp.
		WithStatus(200).WithMessage(string(respJson))
}
