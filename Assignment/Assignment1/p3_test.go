package Assignment1

import (
	"testing"
)

func TestEvenSum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{0, 2, 3, 4}, 6},
		{[]int{-11, 3, 7}, 0},
		{[]int{1, 3, 5, 7}, 0},
		{[]int{-2, -4, -6, -8}, -20},
		{[]int{2, 4, 6, 8}, 20},
	}

	for _, test := range tests {
		if got := EvenSum(test.nums); got != test.want {
			t.Errorf("EvenSum(%v) = %d, want %d", test.nums, got, test.want)
		}
	}
}
