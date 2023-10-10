package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	got := Add(3,3)
	want := 6
	// require = FailNow
	require.Equal(t, want, got, "Add(2,3) should return 5")

	got1 := Add(2,3)
	want1 := 5
	// assert = Fail
	assert.Equal(t, want1, got1, "Add(2,3) should return 5")
}

func TestDis(t *testing.T) {
	tests := []struct {
		product  Product
		expected float64
	}{
		{Product{Name: "Shirt", Price: 100, Discount: 0.1}, 90},
		{Product{Name: "Shoes", Price: 200, Discount: 0.2}, 160},
		{Product{Name: "Cap", Price: 50, Discount: 0}, 50},
	}

	for _, tt := range tests {
		got := tt.product.GetDiscountedPrice()
		assert.Equal(t, tt.expected, got, "For Product %s with Price %.2f and discount %.2f, expected %.2f but got %.2f", tt.product.Name, tt.product.Price, tt.product.Discount, tt.expected, got)
	}
}