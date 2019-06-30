package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-lambda-go/events"
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
	var req EewData
	err := json.Unmarshal([]byte(r.Body), &req)
	if err != nil {
		log.Println(err)
		jsonBytes, _ := json.Marshal(Body{Result: "JSON Unmarshal failed."})
		return events.APIGatewayProxyResponse{StatusCode: 200, Headers: map[string]string{}, Body: string(jsonBytes)}, nil
	}

	file, err := os.Open("./hypo.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	hypo := map[int]string{}

	for {
		line, err := reader.Read()
		if err != nil {
			break
		}
		hypoCode, _ := strconv.Atoi(line[0])
		hypo[hypoCode] = line[1]
	}

	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := dynamodb.New(sess)

	fmt.Println(req)
	fmt.Println(hypo[req.Details.Eewinfo.Hypocode])

	putParams := &dynamodb.PutItemInput{
		TableName: aws.String("Disaster-test"),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(uuid.New().String()),
			},
			"latitude": {
				N: aws.String(fmt.Sprintf("%g", req.Details.Eewinfo.Latitude)),
			},
			"longitude": {
				N: aws.String(fmt.Sprintf("%g", req.Details.Eewinfo.Longitude)),
			},
			"label": {
				S: aws.String(hypo[req.Details.Eewinfo.Hypocode]),
			},
			"magnitude": {
				N: aws.String(fmt.Sprintf("%g", req.Details.Eewinfo.Magnitude)),
			},
			"depth": {
				N: aws.String(fmt.Sprintf("%d", req.Details.Eewinfo.Depth)),
			},
			"time": {
				N: aws.String(fmt.Sprintf("%d", req.Details.Eewinfo.OccuredDatetime)),
			},
		},
	}

	fmt.Println(putParams.Item)

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
