package utils

import "encoding/base64"

func CountEnglishChars(text *string) int {
	var count int = 0
	for _, r := range *text {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			count++
		}
	}
	return count
}

func EncodeBase64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func DecodeBase64(text string) string {
	data, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return ""
	}
	return string(data)
}
