package handler

import (
	"encoding/json"
	"fmt"

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

func NewDataRespHandler(key string, data interface{}) (string, error) {
	jsonData := map[string]interface{}{key: data}
	res, err := json.Marshal(ApiResp{
		Data: jsonData,
	})
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	return string(res), nil
}

func NewErrRespHandler(err error) (string, error) {
	res, err := json.Marshal(ApiResp{
		Data: fmt.Sprintf("error: %v", err.Error()),
	})
	if err != nil {
		return "", utils.NewUnexpectedServerError()
	}
	return string(res), nil
}
