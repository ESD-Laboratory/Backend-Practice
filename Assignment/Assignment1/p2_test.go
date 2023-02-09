package Assignment1

import (
	"testing"
)

func TestCleanUp(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"I am learning Go???", "I am learning Go"},
		{"Everything is good", "Everything is good"},
		{"    Independence day: August 17th, 1945 ", "Independence day August 17th 1945"},
		{"    Go  is  amazing !!! ", "Go is amazing"},
		{"This is a good day     ", "This is a good day"},
		{"   ", ""},
	}
	for _, test := range tests {
		result := CleanUp(test.input)
		if result != test.expected {
			t.Errorf("CleanUp(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}
