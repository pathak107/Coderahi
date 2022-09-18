package handler

import (
	"encoding/json"

	"github.com/pathak107/coderahi-learn/pkg/utils"
)

type ApiResp struct {
	Data interface{} `json:"data"`
}

func NewSuccessCreationRespHandler() (string, error) {
	res, err := json.Marshal(ApiResp{
		Data: "created successfully",
	})
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	return string(res), nil
}

func NewSuccessDeletionRespHandler() (string, error) {
	res, err := json.Marshal(ApiResp{
		Data: "deleted successfully",
	})
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	return string(res), nil
}

func NewSuccessEditRespHandler() (string, error) {
	res, err := json.Marshal(ApiResp{
		Data: "updated successfully",
	})
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	return string(res), nil
}

func NewDataRespHandler(data interface{}) (string, error) {
	res, err := json.Marshal(ApiResp{
		Data: data,
	})
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	return string(res), nil
}
