package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	data := Note{
		Title: "Baca",
		Body:  "lorem of dolor",
	}

	j, _ := json.MarshalIndent(&data, "", " ")

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		StatusCode: 200,
		Body:       string(j),
	}, nil

}

func main() {
	lambda.Start(Handler)
}
