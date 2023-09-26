package main

import "fmt"

type Result struct {
	Operation string
	Value float64
}

func main() {
	chInput := make(chan [2]float64, 4)
	chOutput := make(chan Result, 4)
	
	for i := 0; i < 4; i ++ {
		chInput <- [2]float64{5,3}
	}

	go penjumlahan(chInput, chOutput)
	go pengurangan(chInput, chOutput)
	go perkalian(chInput, chOutput)
	go pembagian(chInput, chOutput)



	for i := 0; i < 4; i ++ {
		result := <- chOutput
		fmt.Printf("%v \t: %.2f\n", result.Operation, result.Value)
	}
}

func penjumlahan(chInput chan [2]float64, chOutput chan Result) {
	data := <- chInput
	chOutput <- Result{"Pertambahan", data[0] + data[1]}
}

func pengurangan(chInput chan [2]float64, chOutput chan Result) {
	data := <- chInput
	chOutput <- Result{"Pengurangan", data[0] - data[1]}
}

func perkalian(chInput chan [2]float64, chOutput chan Result) {
	data := <- chInput
	chOutput <- Result{"Perkalian", data[0] * data[1]}
}

func pembagian(chInput chan [2]float64, chOutput chan Result) {
	data := <- chInput

	if data[1] != 0 {
		chOutput <- Result{"Pembagian", data[0] / data[1]}
	} else {
		chOutput <- Result{"Pembagian", 0.0}
	}
}