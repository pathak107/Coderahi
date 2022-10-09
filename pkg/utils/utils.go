package utils

import "encoding/json"

func ToStringPtr(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

func ToString(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func ToJsonString(key string, data interface{}) string {
	jsonMap := make(map[string]interface{})
	jsonMap[key] = data
	jsonStr, _ := json.Marshal(jsonMap)
	return string(jsonStr)
}
