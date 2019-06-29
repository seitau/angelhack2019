package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-lambda-go/events"

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
func HandleLambdaEvent(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req Request
	err := json.Unmarshal([]byte(r.Body), &req)
	if err != nil {
		log.Println(err)
		jsonBytes, _ := json.Marshal(Body{Result: "JSON Unmarshal failed."})
		return events.APIGatewayProxyResponse{StatusCode: 200, Headers: map[string]string{}, Body: string(jsonBytes)}, nil
	}

	if req.Name == "" || req.EthAddress == "" || req.AdaAddress == "" {
		jsonBytes, _ := json.Marshal(Body{Result: "No data input."})
		return events.APIGatewayProxyResponse{StatusCode: 200, Headers: map[string]string{}, Body: string(jsonBytes)}, nil
	}

	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := dynamodb.New(sess)

	fmt.Println(req)

	queryParams := &dynamodb.QueryInput{
		TableName: aws.String("User-test"),
		ExpressionAttributeNames: map[string]*string{
			"#UserName": aws.String("username"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":username": {
				S: aws.String(req.Name),
			},
		},
		KeyConditionExpression: aws.String("#UserName=:username"),
		IndexName:              aws.String("username-index"),
	}

	queryItem, queryErr := svc.Query(queryParams)
	if queryErr == nil {
		jsonBytes, _ := json.Marshal(Body{Result: "Username already used."})
		return events.APIGatewayProxyResponse{StatusCode: 200, Headers: map[string]string{}, Body: string(jsonBytes)}, nil
	}
	log.Println(queryErr)
	fmt.Println(queryItem)

	putParams := &dynamodb.PutItemInput{
		TableName: aws.String("User-test"),
		Item: map[string]*dynamodb.AttributeValue{
			"ethAddress": {
				S: aws.String(req.EthAddress),
			},
			"adaAddress": {
				S: aws.String(req.AdaAddress),
			},
			"username": {
				S: aws.String(req.Name),
			},
		},
	}

	_, putErr := svc.PutItem(putParams)
	if putErr != nil {
		log.Println(putErr)
		jsonBytes, _ := json.Marshal(Body{Result: "failed"})
		return events.APIGatewayProxyResponse{StatusCode: 200, Headers: map[string]string{}, Body: string(jsonBytes)}, nil
	}

	jsonBytes, _ := json.Marshal(Body{Result: "success"})
	return events.APIGatewayProxyResponse{StatusCode: 200, Headers: map[string]string{}, Body: string(jsonBytes)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
