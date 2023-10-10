package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(2,3)
	want := 7
	
	// if got != want {
	// 	t.Fail()
	// 	t.Logf("Add(2,3) = %v; want %v", got, want)
	// }

	if got != want {
		t.Fatalf("Add(2,3) = %v; want %v", got, want)
	}
}