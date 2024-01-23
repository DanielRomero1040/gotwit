package routers

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"

	"github.com/DanielRomero1040/gotwit/awsgo"
	"github.com/DanielRomero1040/gotwit/db"
	"github.com/DanielRomero1040/gotwit/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

func GetImage(ctx context.Context, typeImage string, request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	resp := models.NewRespApi()

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		return resp.
			WithMessage("El parametro ID es obligatorio ")
	}
	profile, err := db.FindProfile(ID)
	if err != nil {
		return resp.
			WithMessage("Usuario no encontrado " + err.Error())
	}
	var filename string

	switch typeImage {
	case "A":
		filename = profile.Avatar
	case "B":
		filename = profile.Banner
	}
	fmt.Println("Filename " + filename)

	svc := s3.NewFromConfig(awsgo.Cfg)

	file, err := downloadFromS3(ctx, svc, filename)
	if err != nil {
		return resp.
			WithStatus(500).
			WithMessage("Error descargando el archivo s3 " + err.Error())
	}

	customResp := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       file.String(),
		Headers: map[string]string{
			"Content-Type":        "application/octet-stream",
			"Content-Disposition": fmt.Sprintf("attachment; filename=\"%s\"", filename),
		},
	}

	return resp.WithCustomResp(customResp)
}

func downloadFromS3(ctx context.Context, svc *s3.Client, filename string) (*bytes.Buffer, error) {
	bucket := ctx.Value(models.Key("bucketName")).(string)
	obj, err := svc.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, err
	}
	defer obj.Body.Close()
	fmt.Println("Bucket name = " + bucket)

	file, err := ioutil.ReadAll(obj.Body)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(file)
	return buf, nil
}
