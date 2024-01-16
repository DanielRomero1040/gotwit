package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DanielRomero1040/gotwit/db"
	"github.com/DanielRomero1040/gotwit/models"
)

func Register(ctx context.Context) models.RespApi {
	var t models.User
	var res = models.NewRespApi()

	fmt.Println("Entre a Register")
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		res.WithMessage(err.Error())
		fmt.Println(res.GetMessage())
		return res
	}

	if len(t.Email) == 0 {
		res.WithMessage("Debe especificar el email")
		fmt.Println(res.GetMessage())
		return res
	}

	if len(t.Password) < 6 {
		res.WithMessage("Debe ingresar una password de minimo 6 caracteres")
		fmt.Println(res.GetMessage())
		return res
	}

	_, found, _ := db.CheckValidUser(t.Email)
	if found {
		res.WithMessage("ya existe un usuario con ese email ")
		fmt.Println(res.GetMessage())
		return res
	}

	_, status, err := db.InsertResgister(t)
	if err != nil {
		res.WithMessage("ocurrio un error al intentar insertar el registro de usuario " + err.Error())
		fmt.Println(res.GetMessage())
		return res
	}
	if !status {
		res.WithMessage("No se ha logrado insertar elñ registro de usuario ")
		fmt.Println(res.GetMessage())
		return res
	}

	res.WithStatus(200).WithMessage("Registro OK")
	fmt.Println(res.GetMessage())
	return res
}
