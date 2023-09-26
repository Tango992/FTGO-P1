package main

import "fmt"


func main() {
	chInput := make(chan [2]float64, 4)
	chAddOutput := make(chan float64)
	chSubOutput := make(chan float64)
	chMulOutput := make(chan float64)
	chDivOutput := make(chan float64)
	
	for i := 0; i < 4; i ++ {
		chInput <- [2]float64{5,3}
	}

	go penjumlahan(chInput, chAddOutput)
	go pengurangan(chInput, chSubOutput)
	go perkalian(chInput, chMulOutput)
	go pembagian(chInput, chDivOutput)



	for i := 0; i < 4; i ++ {
		select {
		case add := <- chAddOutput:
			fmt.Printf("Penjumlahan\t: %.2f\n", add)
		case sub := <- chSubOutput:
			fmt.Printf("Pengurangan\t: %.2f\n", sub)
		case mul := <- chMulOutput:
			fmt.Printf("Perkalian\t: %.2f\n", mul)
		case div := <- chDivOutput:
			fmt.Printf("Pembagian\t: %.2f\n", div)
		}
	}
}

func penjumlahan(chInput chan [2]float64, chOutput chan float64) {
	data := <- chInput
	chOutput <- data[0] + data[1]
}

func pengurangan(chInput chan [2]float64, chOutput chan float64) {
	data := <- chInput
	chOutput <- data[0] - data[1]
}

func perkalian(chInput chan [2]float64, chOutput chan float64) {
	data := <- chInput
	chOutput <- data[0] * data[1]
}

func pembagian(chInput chan [2]float64, chOutput chan float64) {
	data := <- chInput

	if data[1] != 0 {
		chOutput <- data[0] / data[1]
	} else {
		chOutput <- 0.0
	}
}