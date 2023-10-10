package main

type Product struct {
	Name     string
	Price    float64
	Discount float64
}

func (p *Product) GetDiscountedPrice() float64 {
	return p.Price * (1 - p.Discount)
}

func Add(x, y int) int {
	return x + y
}
