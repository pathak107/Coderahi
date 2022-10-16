package utils

import (
	"encoding/json"
	"path/filepath"
	"runtime"
)

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

func ToIntPtr(num int) *int {
	return &num
}

func ToInt(num *int) int {
	if num == nil {
		return 0
	}
	return *num
}

func ToJsonString(key string, data interface{}) string {
	jsonMap := make(map[string]interface{})
	jsonMap[key] = data
	jsonStr, _ := json.Marshal(jsonMap)
	return string(jsonStr)
}

func getCurrDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func GetImagesDir() string {
	return filepath.Join(getCurrDir(), "../", "../", "public", "images") + "/"
}

func GetPublicDir() string {
	return filepath.Join(getCurrDir(), "../", "../", "public") + "/"
}
