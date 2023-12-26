package gopalindrome_test

import (
	"testing"

	"github.com/aZ4ziL/gopalindrome"
)

func TestReverse(t *testing.T) {
	result := gopalindrome.Reverse("world")
	expected := "dlrow"

	if result != expected {
		t.Errorf("Reverse function returned incorrect result, got: %s, want: %s", result, expected)
	}
}
