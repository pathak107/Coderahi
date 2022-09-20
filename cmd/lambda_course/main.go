package main

import (
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pathak107/coderahi-learn/pkg/handler"
	"github.com/pathak107/coderahi-learn/pkg/services"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//Initital setup
	db, err := services.NewDatabaseService()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	h := handler.NewHandler(db)

	ApiResponse := events.APIGatewayProxyResponse{}
	if strings.Contains(request.Path, "/section") {
		// Switch for identifying the HTTP request
		switch request.HTTPMethod {
		case "POST":
			if strings.Contains(request.Path, "/order") {
				data, err := h.ChangeSectionOrder(request.Body)
				if err != nil {
					return handler.HandleError(err), nil
				}
				ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}
			} else {
				data, err := h.CreateSection(request.Body)
				if err != nil {
					return handler.HandleError(err), nil
				}
				ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}
			}

		case "DELETE":
			data, err := h.DeleteSectionByID(request.PathParameters["section_id"])
			if err != nil {
				return handler.HandleError(err), nil
			}
			ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}

		case "PATCH":
			data, err := h.EditSectionByID(request.Body, request.PathParameters["section_id"])
			if err != nil {
				return handler.HandleError(err), nil
			}
			ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}
		}

	} else if strings.Contains(request.Path, "/subsection") {
		// Switch for identifying the HTTP request
		switch request.HTTPMethod {
		case "POST":
			if strings.Contains(request.Path, "/order") {
				data, err := h.ChangeSubsectionOrder(request.Body)
				if err != nil {
					return handler.HandleError(err), nil
				}
				ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}
			} else {
				data, err := h.CreateSubsection(request.Body)
				if err != nil {
					return handler.HandleError(err), nil
				}
				ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}
			}

		case "DELETE":
			data, err := h.DeleteSubsectionByID(request.PathParameters["subsection_id"])
			if err != nil {
				return handler.HandleError(err), nil
			}
			ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}

		case "PATCH":
			data, err := h.EditSubsectionByID(request.Body, request.PathParameters["subsection_id"])
			if err != nil {
				return handler.HandleError(err), nil
			}
			ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}
		}
	} else {
		// Switch for identifying the HTTP request
		switch request.HTTPMethod {
		case "GET":
			if request.PathParameters["course_id"] != "" {
				data, err := h.FindCourseByID(request.PathParameters["course_id"], request.QueryStringParameters)
				if err != nil {
					return handler.HandleError(err), nil
				}
				ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}
			} else {
				data, err := h.FindAllCourses()
				if err != nil {
					return handler.HandleError(err), nil
				}
				ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}
			}

		case "POST":
			//validates json and returns error if not working
			// err := fastjson.Validate(request.Body)
			data, err := h.CreateCourse(request.Body)
			if err != nil {
				return handler.HandleError(err), nil
			}
			ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}
		case "DELETE":
			data, err := h.DeleteCourse(request.PathParameters["course_id"])
			if err != nil {
				return handler.HandleError(err), nil
			}
			ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}

		case "PATCH":
			data, err := h.EditCourse(request.Body, request.PathParameters["course_id"])
			if err != nil {
				return handler.HandleError(err), nil
			}
			ApiResponse = events.APIGatewayProxyResponse{Body: data, StatusCode: 200}
		}
	}

	// Response
	return ApiResponse, nil
}

func main() {
	lambda.Start(HandleRequest)
}
