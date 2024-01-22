package routers

import (
	"context"
	"encoding/json"

	"github.com/DanielRomero1040/gotwit/db"
	"github.com/DanielRomero1040/gotwit/models"
)

func UpdateProfile(ctx context.Context, claim models.Claim) models.RespApi {
	resp := models.NewRespApi()
	var user models.User

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &user)
	if err != nil {
		return resp.WithMessage("Datos incorrectos " + err.Error())
	}
	status, err := db.UpdateRegister(user, claim.ID.Hex())
	if err != nil {
		return resp.WithMessage("ocurrio un error al intentar modificar el registro " + err.Error())
	}

	if !status {
		return resp.WithMessage("No se ha logrado modificar el registro ")
	}
	return resp.
		WithStatus(200).
		WithMessage("Modificacion de profile OK ")
}
