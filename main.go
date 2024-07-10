package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/DanielRomero1040/gotwit/awsgo"
	"github.com/DanielRomero1040/gotwit/db"
	"github.com/DanielRomero1040/gotwit/handlers"
	"github.com/DanielRomero1040/gotwit/models"
	"github.com/DanielRomero1040/gotwit/secretmanager"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	awsgo.InicializarAWS()

	//valido variables de entorno
	if !EnvVariablesValidation() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la variables de entorno. Deben incluir 'SecretName', 'BucketName' y 'UrlPrefix' ",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura de secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	path := strings.Replace(request.PathParameters["gotwit"], os.Getenv("UrlPrefix"), "", -1) //mirar en la apigateway que sea el mismo prefix

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtsign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	//chequeo conexion base datos - conecto base de datos
	err = db.ConectionDB(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error conectando en la base de datos " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}
	//handlers / controller
	respApi := handlers.Handlers(awsgo.Ctx, request)

	if respApi.GetCustomResp() == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: respApi.GetStatus(),
			Body:       respApi.GetMessage(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	} else {
		return respApi.GetCustomResp(), nil
	}
}

func EnvVariablesValidation() bool {
	fmt.Println("Cargando variables de entorno ")
	_, envVarExist := os.LookupEnv("SecretName")
	if !envVarExist {
		return envVarExist
	}
	_, envVarExist = os.LookupEnv("BucketName")
	if !envVarExist {
		return envVarExist
	}
	_, envVarExist = os.LookupEnv("UrlPrefix")

	return envVarExist
}
