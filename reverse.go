package gopalindrome

// Reverse is function to reverse a text.
func Reverse(text string) string {
	runes := []rune(text)
	result := make([]rune, len(runes))

	for i, j := len(runes)-1, 0; i >= 0 && j < len(runes); i, j = i-1, j+1 {
		result[j] = runes[i]
	}

	return string(result)
}
