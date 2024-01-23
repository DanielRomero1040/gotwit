package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/DanielRomero1040/gotwit/awsgo"
	"github.com/DanielRomero1040/gotwit/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error) {
	var datosSecret models.Secret
	fmt.Println("-> Pido secreto" + secretName)
	smc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := smc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err

	}
	err = json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println("-> Lectura de Secret OK " + secretName)
	return datosSecret, err
}
