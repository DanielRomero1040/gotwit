package routers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/DanielRomero1040/gotwit/db"
	"github.com/DanielRomero1040/gotwit/models"
)

func SaveTweet(ctx context.Context, claim models.Claim) models.RespApi {
	var message models.Tweet
	res := models.NewRespApi()
	IDUsuario := claim.ID.Hex()
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &message)
	if err != nil {
		res.WithMessage("Ocurrio un error al intentar decodificar el body " + err.Error())
		return res
	}

	registro := models.SaveTweet{
		UserID:  IDUsuario,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.InsertTweet(registro)
	if err != nil {
		res.WithMessage("Ocurrio un error al intentar insertar el registro " + err.Error())
		return res
	}

	if !status {
		res.WithMessage("No se ha logrado insertar el Tweet ")
		return res
	}

	return res.WithStatus(200).WithMessage("Tweet Creado Correctamente ")
}
