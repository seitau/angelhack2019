package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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

// Change ...
type Change struct {
	ID        string  `json:"id"`
	Depth     int     `json:"depth"`
	Label     string  `json:"label"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Magnitude float64 `json:"magnitude"`
	Time      int     `json:"time"`
}

// User ...
type User struct {
	UserName   string
	EthAddress string
	AdaAddress string
}

// DynamoEventChange ...
type DynamoEventChange struct {
	NewImage *dynamodb.AttributeValue `json:"NewImage"`
}

// DynamoEventRecord ...
type DynamoEventRecord struct {
	Change    DynamoEventChange `json:"dynamodb"`
	EventName string            `json:"eventName"`
	EventID   string            `json:"eventID"`
}

// DynamoEvent ...
type DynamoEvent struct {
	Records []DynamoEventRecord `json:"records"`
}

// HandleLambdaEvent ...
func HandleLambdaEvent(event events.DynamoDBEvent) error {
	var newRecord Change
	for _, record := range event.Records {
		if record.EventName == "INSERT" {
			id := record.Change.NewImage["id"].String()
			depth, err := record.Change.NewImage["depth"].Integer()
			if err != nil {
				return err
			}
			label := record.Change.NewImage["label"].String()
			latitude, err := record.Change.NewImage["latitude"].Float()
			if err != nil {
				return err
			}
			longitude, err := record.Change.NewImage["longitude"].Float()
			if err != nil {
				return err
			}
			magnitude, err := record.Change.NewImage["magnitude"].Float()
			if err != nil {
				return err
			}
			time, err := record.Change.NewImage["time"].Integer()
			if err != nil {
				return err
			}

			newRecord := Change{
				ID:        id,
				Depth:     int(depth),
				Label:     label,
				Latitude:  latitude,
				Longitude: longitude,
				Magnitude: magnitude,
				Time:      int(time),
			}
			fmt.Println(newRecord)
		}
	}

	area := detectArea(newRecord.Latitude, newRecord.Longitude)

	queryParams := &dynamodb.QueryInput{
		TableName: aws.String("User-test"),
		ExpressionAttributeNames: map[string]*string{
			"#Latitude": aws.String("recent_latitude"),
			//"#Longitude": aws.String("recent_longitude"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":low_latitude": {
				N: aws.String(fmt.Sprintf("%g", area[0])),
			},
			":high_latitude": {
				N: aws.String(fmt.Sprintf("%g", area[1])),
			},
			//":low_longitude": {
			//	N: aws.String(fmt.Sprintf("%g", area[2])),
			//},
			//":high_longitude": {
			//	N: aws.String(fmt.Sprintf("%g", area[3])),
			//},
		},
		KeyConditionExpression: aws.String("#Latitude >= :low_latitude AND #Latitude <= :high_latitude"),
		IndexName:              aws.String("recent_latitude-recent_longitude-index"),
	}

	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := dynamodb.New(sess)

	queryItem, queryErr := svc.Query(queryParams)
	if queryErr != nil {
		log.Println(queryErr)
		return nil
	}
	users := []User{}
	fmt.Println("イェイ")
	fmt.Println(queryItem)

	if err := dynamodbattribute.UnmarshalListOfMaps(queryItem.Items, &users); err != nil {
		log.Println(queryErr)
		return nil
	}

	return nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
