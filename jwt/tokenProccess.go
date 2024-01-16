package jwt

import (
	"errors"
	"strings"

	"github.com/DanielRomero1040/gotwit/models"
	"github.com/golang-jwt/jwt/v5"
)

var (
	Email     string
	IDUsuario string
)

func TokenProccess(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	miClave := []byte(JWTSign)
	var claims models.Claim

	splittedToken := strings.Split(tk, "Bearer")
	if len(splittedToken) != 2 {
		return &claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splittedToken[1])
	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		//rutina checkea contra db
	}
	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Token Invalido")
	}

	return &claims, false, string(""), err
}
