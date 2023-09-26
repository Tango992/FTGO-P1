package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Diameter int
	Area          float64
	Circumference float64
}

func main() {
	diameters := []int{7, 14, 30, 46}
	ch := make(chan Circle, 4)
	defer close(ch)

	results := []Circle{}
	for _, diameter := range diameters {
		go calculateCircle(diameter, ch)
		results = append(results, <-ch)
	}
	
	for _, result := range results {
		fmt.Printf("\nDiameter\t: %vcm\n", result.Diameter)
		fmt.Printf("Circumference\t: %.2fcm²\n", result.Circumference)
		fmt.Printf("Area\t\t: %.2fcm³\n", result.Area)
	}
}

func calculateCircle(d int, ch chan Circle) {
	area := math.Pi * float64(d/2) * float64(d/2)
	circumference := math.Pi * float64(d)
	ch <- Circle{d, area, circumference}
}