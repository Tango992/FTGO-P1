package main

import "testing"

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
		if got != tt.expected {
			t.Logf("For Product %s with Price %f and discount %f, expected %f but got %f", tt.product.Name, tt.product.Price, tt.product.Discount, tt.expected, got)
			t.Fail()
		}
	}
}
