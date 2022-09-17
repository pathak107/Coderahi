package handler

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/pathak107/coderahi-learn/pkg/utils"
)

func HandleError(err error) events.APIGatewayProxyResponse {
	e, ok := err.(*utils.ApiError)
	if ok {
		return events.APIGatewayProxyResponse{
			StatusCode: e.StatusCode(),
			Body:       e.Err.Error(),
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       "some unexpected error",
	}
}
