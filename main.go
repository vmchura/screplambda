package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-lambda-go/events"
	"github.com/icza/screp/repparser"
	"log"
)
type RequestReplay struct{
	Value string `json:"replayfile"`
	FileName string `json:"filename"`
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))
	log.Println("Start")

	data := &RequestReplay{

	}
	err := json.Unmarshal([]byte(request.Body), data)

	if err != nil{
		return events.APIGatewayProxyResponse { Body: "", StatusCode: 500 }, err
	}

	log.Println("FileName")
	log.Println(fmt.Sprintln(data.FileName))
	log.Println("End File Name")
	p, err := base64.StdEncoding.DecodeString(data.Value)
	if err != nil {
		return events.APIGatewayProxyResponse { Body: "", StatusCode: 500 }, err
	}
	log.Println("Descode passes")
	r, err := repparser.Parse(p)
	r.Compute()
	r.Commands = nil
	r.MapData = nil

	if err != nil {
		return events.APIGatewayProxyResponse { Body: "", StatusCode: 500 }, err
	}
	log.Println("parse passes")

	res,err := json.Marshal(r)
	if err != nil {
		return events.APIGatewayProxyResponse { Body: "", StatusCode: 500 }, err
	}

	return events.APIGatewayProxyResponse { Body: string(res), StatusCode: 200 }, nil



}



func main() {
	lambda.Start(handleRequest)

}

