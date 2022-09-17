package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pathak107/coderahi-learn/pkg/handler"
	"github.com/pathak107/coderahi-learn/pkg/services"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := services.NewDatabaseService()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	h := handler.NewHandler(db)

	ApiResponse := events.APIGatewayProxyResponse{}
	// Switch for identifying the HTTP request
	switch request.HTTPMethod {
	case "GET":
		if request.PathParameters["post_id"] != "" {
			h.FindPostByID(request.PathParameters["post_id"])
			ApiResponse = events.APIGatewayProxyResponse{Body: "Hey " + name + " welcome! ", StatusCode: 200}
		} else {
			posts, err := h.FindAllPosts()
			if err != nil {
				return handler.HandleError(err), nil
			}
			ApiResponse = events.APIGatewayProxyResponse{Body: posts, StatusCode: 200}
		}

	case "POST":
		//validates json and returns error if not working
		err := fastjson.Validate(request.Body)

		if err != nil {
			body := "Error: Invalid JSON payload ||| " + fmt.Sprint(err) + " Body Obtained" + "||||" + request.Body
			ApiResponse = events.APIGatewayProxyResponse{Body: body, StatusCode: 500}
		} else {
			ApiResponse = events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}
		}

	}
	// Response
	return ApiResponse, nil
}

func main() {
	lambda.Start(HandleRequest)
}
