package db

import "golang.org/x/crypto/bcrypt"

func EncrypPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	if err != nil {
		return err.Error(), err
	}
	return string(bytes), nil
}
