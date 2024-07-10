package routers

import (
	"encoding/json"

	"github.com/DanielRomero1040/gotwit/db"
	"github.com/DanielRomero1040/gotwit/models"
	"github.com/aws/aws-lambda-go/events"
)

func GetRelation(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	resp := models.NewRespApi()

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		return *resp.
			WithMessage("El parametro ID es obligatorio ")
	}
	var relation models.Relation
	relation.UserID = claim.ID.Hex()
	relation.UserRelationId = ID

	var respRelation models.RespRelation

	isRelation := db.GetRelation(relation)

	respRelation.Status = isRelation

	respJson, err := json.Marshal(respRelation)
	if err != nil {
		return *resp.
			WithStatus(500).
			WithMessage("Error al formatear los datos de los usuarios como JSON " + err.Error())
	}
	return *resp.
		WithStatus(200).
		WithMessage(string(respJson))

}
