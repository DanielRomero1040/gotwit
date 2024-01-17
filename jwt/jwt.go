package jwt

import (
	"context"
	"time"

	"github.com/DanielRomero1040/gotwit/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(ctx context.Context, user models.User) (string, error) {
	jwtSign := ctx.Value(models.Key("jwtsign")).(string)
	myKey := []byte(jwtSign)

	payload := jwt.MapClaims{
		"email":            user.Email,
		"nombre":           user.Name,
		"apellido":         user.Lastname,
		"fecha_nacimiento": user.Birthday,
		"biografia":        user.Biography,
		"ubicacion":        user.Location,
		"sitioweb":         user.Website,
		"_id":              user.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
