package main

import "fmt"

func main() {
	chAddinput := make(chan [2]float64)
	chSubinput := make(chan [2]float64)
	chMulinput := make(chan [2]float64)
	chDivinput := make(chan [2]float64)
	chAddOutput := make(chan float64)
	chSubOutput := make(chan float64)
	chMulOutput := make(chan float64)
	chDivOutput := make(chan float64)

	go penjumlahan(chAddinput, chAddOutput)
	go pengurangan(chSubinput, chSubOutput)
	go perkalian(chMulinput, chMulOutput)
	go pembagian(chDivinput, chDivOutput)

	chAddinput <- [2]float64{5,3}
	chSubinput <- [2]float64{5,3}
	chMulinput <- [2]float64{5,3}
	chDivinput <- [2]float64{5,3}

	fmt.Printf("Penjumlahan\t: %.2f\n", <- chAddOutput)
	fmt.Printf("Pengurangan\t: %.2f\n", <- chSubOutput)
	fmt.Printf("Perkalian\t: %.2f\n", <- chMulOutput)
	fmt.Printf("Pembagian\t: %.2f\n", <- chDivOutput)
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
	}
	chOutput <- 0.0
}