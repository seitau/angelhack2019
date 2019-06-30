package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-lambda-go/lambda"
)

// Request ...
type Request struct {
	Name      string  `json:"username"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// User ...
type User struct {
	UserName   string
	EthAddress string
	AdaAddress string
}

// Body ...
type Body struct {
	Result string `json:"result"`
}

// HandleLambdaEvent ...
func HandleLambdaEvent(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var req Request
	fmt.Println(r.Body)
	err := json.Unmarshal([]byte(r.Body), &req)
	if err != nil {
		log.Println(err)
		jsonBytes, _ := json.Marshal(Body{Result: "JSON Unmarshal failed."})
		return events.APIGatewayProxyResponse{StatusCode: 200, Headers: map[string]string{}, Body: string(jsonBytes)}, nil
	}

	if req.Name == "" || req.Latitude < -90 || 90 < req.Latitude || req.Longitude < -180 || 180 < req.Longitude {
		jsonBytes, _ := json.Marshal(Body{Result: "No data input."})
		return events.APIGatewayProxyResponse{StatusCode: 200, Headers: map[string]string{}, Body: string(jsonBytes)}, nil
	}

	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := dynamodb.New(sess)

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
	if queryErr != nil {
		log.Println(queryErr)
		jsonBytes, _ := json.Marshal(Body{Result: "Username not found."})
		return events.APIGatewayProxyResponse{StatusCode: 200, Headers: map[string]string{}, Body: string(jsonBytes)}, nil
	}
	users := []User{}

	if err := dynamodbattribute.UnmarshalListOfMaps(queryItem.Items, &users); err != nil {
		log.Println(queryErr)
		jsonBytes, _ := json.Marshal(Body{Result: fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err)})
		return events.APIGatewayProxyResponse{StatusCode: 200, Headers: map[string]string{}, Body: string(jsonBytes)}, nil
	}
	//fmt.Println(users[0])

	location := &dynamodb.AttributeValue{
		M: map[string]*dynamodb.AttributeValue{
			"timestamp": {N: aws.String(fmt.Sprintf("%f", float64(time.Now().Unix())))},
			"longitude": {N: aws.String(fmt.Sprintf("%f", req.Longitude))},
			"latitude":  {N: aws.String(fmt.Sprintf("%f", req.Latitude))},
		},
	}
	var locations []*dynamodb.AttributeValue
	locations = append(locations, location)

	updateParams := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ethAddress": {
				S: aws.String(users[0].EthAddress),
			},
			"adaAddress": {
				S: aws.String(users[0].AdaAddress),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":timestamp": location.M["timestamp"],
			":longitude": location.M["longitude"],
			":latitude":  location.M["latitude"],
			":locations": {
				L: locations,
			},
			":empty_list": {
				L: []*dynamodb.AttributeValue{},
			},
		},
		ReturnValues:     aws.String("ALL_NEW"),
		UpdateExpression: aws.String("SET location_history = list_append(if_not_exists(location_history, :empty_list), :locations), recent_timestamp = :timestamp, recent_longitude = :longitude, recent_latitude = :latitude"),
		TableName:        aws.String("User-test"),
	}

	_, updateErr := svc.UpdateItem(updateParams)
	if updateErr != nil {
		log.Println(updateErr)
		jsonBytes, _ := json.Marshal(Body{Result: "failed"})
		return events.APIGatewayProxyResponse{StatusCode: 200, Headers: map[string]string{}, Body: string(jsonBytes)}, nil
	}

	jsonBytes, _ := json.Marshal(Body{Result: "success"})
	return events.APIGatewayProxyResponse{StatusCode: 200, Headers: map[string]string{}, Body: string(jsonBytes)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
