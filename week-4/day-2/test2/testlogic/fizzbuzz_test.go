package fizzbuzz

import "testing"

func TestFizzbuzz(t *testing.T) {
	tests := []struct {
		input  int
		expected string
	}{
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{13, "13"},
	}

	for _, test := range tests {
		got := FizzBuzz(test.input)
		if got != test.expected {
			t.Errorf("FizzBuzz %v = %v; want %v",test.input, got, test.expected)
		}
	}
}
