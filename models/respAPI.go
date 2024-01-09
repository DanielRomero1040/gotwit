package models

import "github.com/aws/aws-lambda-go/events"

type RespApi struct {
	status     int
	message    string
	customResp *events.APIGatewayProxyResponse
}

func (r RespApi) WithStatus(status int) RespApi {
	r.status = status
	return r
}

func (r RespApi) WithMessage(message string) RespApi {
	r.message = message
	return r
}

func (r RespApi) WithCustomResp(customResp *events.APIGatewayProxyResponse) RespApi {
	r.customResp = customResp
	return r
}

func (r RespApi) GetCustomResp() *events.APIGatewayProxyResponse {
	return r.customResp
}

func (r RespApi) GetStatus() int {
	return r.status
}

func (r RespApi) GetMessage() string {
	return r.message
}

//NewRespApi creates a new RespApi struct and sets status 400 and message error by default
func NewRespApi() RespApi {
	return RespApi{
		status:  400,
		message: "Error en conexion con API",
	}
}
