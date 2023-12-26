package gopalindrome_test

import (
	"testing"

	"github.com/aZ4ziL/gopalindrome"
)

func TestPalindrome(t *testing.T) {
	result := gopalindrome.Palindrome("radar")
	expected := true

	if result != expected {
		t.Error("Input: `radar`. The output should be true")
	}
}
