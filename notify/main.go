package main

import (
	"encoding/json"
	"log"
	"strings"
	"net/url"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

type SlackMessage struct {
	Text string `json:"text"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var keyword string
	for _, value := range strings.Split(request.Body, "&") {
		param := strings.Split(value, "=")
		if param[0] == "trigger_word" {
			keyword, _ = url.QueryUnescape(param[1])
		}
	}

	var text string
	if keyword == "hoge" {
		text = "hogeさんに電話がきたよ！"
	} else if keyword == "fuga" {
		text = "fugaさんに電話がきたよ！"
	}
	
	j, err := json.Marshal(SlackMessage{Text: text})	
	if err != nil {
		log.Print(err)
		return events.APIGatewayProxyResponse{Body: "Error"}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(j),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
