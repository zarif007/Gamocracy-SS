package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/zarif007/Gamocracy-SS/Post/pkg/handlers"
)

var (
	dynaClient dynamodbiface.DynamoDBAPI
)

func main() {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)})

	if err != nil {
		return
	}

	dynaClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}

const tableName = "GC_Post"

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetPost(req, tableName, dynaClient)
	case "POST":
		return handlers.CreatePost(req, tableName, dynaClient)
	case "PUT":
		return handlers.UpdatePost(req, tableName, dynaClient)
	case "DELETE":
		return handlers.DeletePost(req, tableName, dynaClient)
	default:
		return handlers.UnhandledMethod()
	}
}