package gopalindrome

// Palindrome a function to check whether a word is a palindrome or not.
func Palindrome(text string) bool {
	reverseText := Reverse(text)

	if reverseText == text {
		return true
	} else {
		return false
	}
}
