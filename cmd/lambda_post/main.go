package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pathak107/coderahi-learn/pkg/auth"
	"github.com/pathak107/coderahi-learn/pkg/handler"
	"github.com/pathak107/coderahi-learn/pkg/services"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db, err := services.NewDatabaseService()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	h := handler.NewHandler(db)

	authService, err := auth.NewJWTService(os.Getenv("JKS_URL"))
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	ApiResponse := events.APIGatewayProxyResponse{}
	// Switch for identifying the HTTP request
	switch request.HTTPMethod {
	case "GET":
		authService.VerifyJWTToken(request.Headers["Authorization"])
		if request.PathParameters["post_id"] != "" {
			data, err := h.FindPostByID(request.PathParameters["post_id"])
			if err != nil {
				return handler.HandleError(err), nil
			}
			ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}
		} else {
			data, err := h.FindAllPosts()
			if err != nil {
				return handler.HandleError(err), nil
			}
			ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}
		}

	case "POST":
		//validates json and returns error if not working
		// err := fastjson.Validate(request.Body)
		data, err := h.CreatePost(request.Body)
		if err != nil {
			return handler.HandleError(err), nil
		}
		ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}
	case "DELETE":
		data, err := h.DeletePost(request.PathParameters["post_id"])
		if err != nil {
			return handler.HandleError(err), nil
		}
		ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}

	case "PATCH":
		data, err := h.EditPost(request.Body, request.PathParameters["post_id"])
		if err != nil {
			return handler.HandleError(err), nil
		}
		ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}
	}
	// Response
	return ApiResponse, nil
}

func main() {
	lambda.Start(HandleRequest)
}
