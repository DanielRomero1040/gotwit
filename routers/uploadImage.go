package routers

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"mime"
	"mime/multipart"
	"strings"

	"github.com/DanielRomero1040/gotwit/db"
	"github.com/DanielRomero1040/gotwit/models"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type readSeeker struct {
	io.Reader
}

func (rs *readSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func UploadImage(ctx context.Context, typeImage string, request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	res := models.NewRespApi()
	IDUser := claim.ID.Hex()
	var filename string
	var user models.User
	bucket := aws.String(ctx.Value(models.Key("bucketName")).(string))

	switch typeImage {
	case "A":
		filename = "avatars/" + IDUser + ".jpg"
		user.Avatar = filename
	case "B":
		filename = "banners/" + IDUser + ".jpg"
		user.Banner = filename
	}

	mediaType, params, err := mime.ParseMediaType(request.Headers["Content-Type"])
	if err != nil {
		return res.
			WithStatus(500).
			WithMessage(err.Error())
	}

	if !strings.HasPrefix(mediaType, "multipart/") {
		return res.
			WithMessage("Debe enviar una imagen con el 'content-Type' de tipo 'multipart/' en el header ").
			WithStatus(400)
	}

	body, err := base64.StdEncoding.DecodeString(request.Body)
	if err != nil {
		return res.
			WithStatus(500).
			WithMessage(err.Error())
	}

	multipartReader := multipart.NewReader(bytes.NewReader(body), params["boundary"])
	p, err := multipartReader.NextPart()

	if err != nil && err != io.EOF {
		return res.
			WithStatus(500).
			WithMessage(err.Error())
	}

	if err != io.EOF {
		if p.FileName() != "" {
			buf := bytes.NewBuffer(nil)
			if _, err := io.Copy(buf, p); err != nil {
				return res.
					WithStatus(500).
					WithMessage(err.Error())
			}
			sess, err := session.NewSession(&aws.Config{
				Region: aws.String("us-east-1"),
			})

			if err != nil {
				return res.
					WithStatus(500).
					WithMessage(err.Error())
			}

			uploader := s3manager.NewUploader(sess)
			_, err = uploader.Upload(&s3manager.UploadInput{
				Bucket: bucket,
				Key:    aws.String(filename),
				Body:   &readSeeker{buf},
			})
			if err != nil {
				return res.
					WithStatus(500).
					WithMessage(err.Error())
			}
		}
	}
	status, err := db.UpdateRegister(user, IDUser)
	if err != nil || !status {
		return res.
			WithStatus(400).
			WithMessage("Error al modificar registro del usuario " + err.Error())
	}

	return res.
		WithStatus(200).
		WithMessage("Imagen uploaded ok! ")
}
