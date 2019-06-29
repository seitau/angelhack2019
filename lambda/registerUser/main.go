package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-lambda-go/lambda"
)

// Request ...
type Request struct {
	Name       string `json:"username"`
	EthAddress string `json:"eth_address"`
	AdaAddress string `json:"ada_address"`
}

// Body ...
type Body struct {
	Result string `json:"result"`
}

// HandleLambdaEvent ...
func HandleLambdaEvent(r Request) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Function Start")
	if r.Name == "" || r.EthAddress == "" || r.AdaAddress == "" {
		jsonBytes, _ := json.Marshal(Body{Result: "No data input."})
		return events.APIGatewayProxyResponse{StatusCode: 200, Headers: map[string]string{}, Body: string(jsonBytes)}, nil
	}
	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := dynamodb.New(sess)

	getParams := &dynamodb.GetItemInput{
		TableName: aws.String("Users-test"),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(r.Name),
			},
		},
	}

	getItem, getErr := svc.GetItem(getParams)
	if getErr == nil {
		jsonBytes, _ := json.Marshal(Body{Result: "Username already used."})
		return events.APIGatewayProxyResponse{StatusCode: 200, Headers: map[string]string{}, Body: string(jsonBytes)}, nil
	}
	fmt.Println(getItem)

	putParams := &dynamodb.PutItemInput{
		TableName: aws.String("Users-test"),
		Item: map[string]*dynamodb.AttributeValue{
			"ethAddress": {
				S: aws.String(r.EthAddress),
			},
			"adaAddress": {
				S: aws.String(r.AdaAddress),
			},
			"username": {
				S: aws.String(r.Name),
			},
		},
	}

	_, putErr := svc.PutItem(putParams)
	if putErr != nil {
		jsonBytes, _ := json.Marshal(Body{Result: "failed"})
		return events.APIGatewayProxyResponse{StatusCode: 200, Headers: map[string]string{}, Body: string(jsonBytes)}, nil
	}
	jsonBytes, _ := json.Marshal(Body{Result: "success"})
	return events.APIGatewayProxyResponse{StatusCode: 200, Headers: map[string]string{}, Body: string(jsonBytes)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
