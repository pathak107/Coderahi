package utils

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
