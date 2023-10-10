package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFizzbuzz(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{13, "13"},
	}
	for _, test := range tests {
		got := FizzBuzz(test.input)
		assert.Equal(t, test.expected, got, "FizzBuzz(%v) got = %v; want %v", test.input, got, test.expected)
	}
}
