package main

import (
	"testing"
)

func TestSumOddNumbers(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 3, 4, 5}, 9},
		{[]int{0, 2, 4, 6, 8}, 0},
		{[]int{1, 3, 5, 7, 9}, 25},
	}

	for _, test := range tests {
		result := SumOddNumbers(test.input)
		if result != test.output {
			t.Errorf("Dengan input %v, harapannya mendapatkan %d, tapi mendapatkan %d", test.input, test.output, result)
		}
	}
}
