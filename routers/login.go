package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/DanielRomero1040/gotwit/db"
	"github.com/DanielRomero1040/gotwit/jwt"
	"github.com/DanielRomero1040/gotwit/models"
	"github.com/aws/aws-lambda-go/events"
)

func Login(ctx context.Context) models.RespApi {
	var user models.User
	resp := models.NewRespApi()

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &user)
	if err != nil {
		resp.WithMessage("Usuario / contraseña Invalidas " + err.Error())
		return resp
	}
	//le puedo agregar otras validaciones como @ o .
	if len(user.Email) == 0 {
		resp.WithMessage("Email usuario es requerido " + err.Error())
		return resp
	}
	userData, exist := db.TryLogin(user.Email, user.Password)
	if !exist {
		resp.WithMessage("Usuario / contraseña Invalidas ")
		return resp
	}
	jwtKey, err := jwt.GenerateJWT(ctx, userData)
	if err != nil {
		resp.WithMessage("ocurrio un error al intentar generar el token correspondiente  " + err.Error())
		return resp
	}
	respLogin := models.ResponseLogin{
		Token: jwtKey,
	}

	token, err2 := json.Marshal(respLogin)
	if err2 != nil {
		resp.WithMessage("ocurrio un error al intentar formatear el token a json  " + err2.Error())
		return resp
	}
	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	}

	cookieString := cookie.String()

	resGateway := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  cookieString,
		},
	}
	return resp.
		WithStatus(200).
		WithMessage(string(token)).
		WithCustomResp(resGateway)
}
