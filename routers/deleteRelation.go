package routers

import (
	"github.com/DanielRomero1040/gotwit/db"
	"github.com/DanielRomero1040/gotwit/models"
	"github.com/aws/aws-lambda-go/events"
)

func DeleteRelation(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	resp := models.NewRespApi()

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		return *resp.
			WithMessage("El parametro ID es obligatorio ")
	}
	var relation models.Relation
	relation.UserID = claim.ID.Hex()
	relation.UserRelationId = ID

	//
	status, err := db.DeleteRelation(relation)

	if err != nil {
		return *resp.
			WithMessage("Ocurrio un error al intentar eliminar la relacion " + err.Error())
	}

	if !status {
		return *resp.
			WithMessage("No se ha logrado eliminar la relacion ")
	}

	return *resp.
		WithStatus(200).
		WithMessage("Baja Relacion OK ")
}
