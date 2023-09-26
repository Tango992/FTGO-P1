package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	chAddinput := make(chan [2]float64)
	chSubinput := make(chan [2]float64)
	chOutput := make(chan float64)

	go penjumlahanUnbuffered(chAddinput, chOutput)
	go penguranganUnbuffered(chSubinput, chOutput)

	chAddinput <- [2]float64{5,3}
	chSubinput <- [2]float64{5,3}

	fmt.Printf("Penjumlahan\t: %.2f\n", <- chOutput)
	fmt.Printf("Pengurangan\t: %.2f\n", <- chOutput)
	fmt.Printf("Elapsed time: %v\n", time.Since(now))
}

func penjumlahanUnbuffered(chInput <-chan [2]float64, chOutput chan<- float64) {
	data := <- chInput
	chOutput <- data[0] + data[1]
}

func penguranganUnbuffered(chInput <-chan [2]float64, chOutput chan<- float64) {
	data := <- chInput
	chOutput <- data[0] - data[1]
}