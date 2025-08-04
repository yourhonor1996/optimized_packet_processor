package utils

func CountEnglishChars(text *string) int {
	var count int = 0
	for _, r := range *text {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			count++
		}
	}
	return count
}
