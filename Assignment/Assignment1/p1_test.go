package Assignment1

import (
	"testing"
)

func TestDivideWhole(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{13, 3, 4},
		{4, 9, 0},
		{7, 0, -1},
		{10, 5, 2},
		{100, 50, 2},
		{11, 7, 1},
	}
	for _, test := range tests {
		result := DivideWhole(test.a, test.b)
		if result != test.expected {
			t.Errorf("DivideWhole(%d, %d) = %d, expected %d", test.a, test.b, result, test.expected)
		}
	}
}
